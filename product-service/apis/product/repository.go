package product

import (
	"encoding/json"
	"khanhnguyen234/product-service/_elastic"
	"khanhnguyen234/product-service/_mongo"
	"khanhnguyen234/product-service/_rabbitmq"
	"khanhnguyen234/product-service/_redis"
	"khanhnguyen234/product-service/common"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "product"
	eIndex     = "product"
	eType      = "byName"
)

func (s *ProductCreate) CreateProductMongo(product ProductCreate) (ProductCreate, error) {
	if product.Id == "" {
		product.Id = uuid.NewV4().String()
	}
	err := _mongo.InsertOne(collection, product)

	return product, err
}

func (s *ProductCreate) CreateProductPub(product ProductCreate) {
	jsonValue, _ := json.Marshal(product)

	e := _rabbitmq.Exchange{
		Exchange:   "product",
		Type:       _rabbitmq.ExchangeDirect,
		RoutingKey: "create",
	}
	e.Pub(string(jsonValue))
}

func CreateProductSub() {
	e := _rabbitmq.Exchange{
		Exchange:   "product",
		Type:       _rabbitmq.ExchangeDirect,
		RoutingKey: "create",
		Queue:      "create",
	}

	msgs := e.Sub()

	go func() {
		for d := range msgs {
			var data ProductCreate
			json.Unmarshal([]byte(d.Body), &data)
			data.CreateProductElastic(data)
			data.CreateProductRedis(data)
		}
	}()
}

func (s *ProductCreate) CreateProductElastic(product ProductCreate) error {
	err := _elastic.Put(eIndex, eType, product.Id, product)
	return err
}

func (s *ProductCreate) CreateProductRedis(product ProductCreate) {
	key := "product_" + product.Id
	redis := _redis.GetRedis()
	jsonValue, _ := json.Marshal(product)
	redis.Set(key, string(jsonValue), 0)
}

func (s *ProductCreate) GetProductsMongo() ([]ProductCreate, error) {
	db := _mongo.ConnectMongo()
	c := db.Collection(collection)

	ctx := common.GetContext()
	condition := bson.D{}
	cur, err := c.Find(ctx, condition)

	var products []ProductCreate

	for cur.Next(ctx) {
		var t ProductCreate
		err := cur.Decode(&t)
		if err != nil {
			return products, err
		}

		products = append(products, t)
	}

	if err := cur.Err(); err != nil {
		return products, err
	}

	return products, err
}

func (s *ProductCreate) GetProductsFlashSaleMongo(query ProductFlashSale) ([]ProductCreate, error) {
	db := _mongo.ConnectMongo()
	c := db.Collection(collection)
	ctx := common.GetContext()

	// 1601448417
	unix := query.Time
	condition := bson.M{
		"flash_sale":            bson.D{{"$eq", true}},
		"flash_sale_unix_start": bson.D{{"$lte", unix}},
		"flash_sale_unix_end":   bson.D{{"$gte", unix}},
	}

	options := options.Find()

	// Sort by `_id` field descending
	options.SetSort(bson.D{{"price", -1}})

	if query.Limit != 0 {
		options.SetLimit(query.Limit)
	}

	cur, err := c.Find(ctx, condition, options)

	var products []ProductCreate

	for cur.Next(ctx) {
		var t ProductCreate
		err := cur.Decode(&t)
		if err != nil {
			return products, err
		}

		redis := _redis.GetRedis()
		jsonValue, _ := json.Marshal(t)
		redis.RPush(strings.Join([]string{"flash_sale", strconv.Itoa(unix)}, "-"), string(jsonValue), 0)

		products = append(products, t)
	}

	if err := cur.Err(); err != nil {
		return products, err
	}

	return products, err
}

func (s *ProductCreate) GetProductsFlashSaleRedis(query ProductFlashSale) ([]ProductCreate, error) {
	var products []ProductCreate
	unix := query.Time

	redis := _redis.GetRedis()
	slices, _ := redis.LRange(strings.Join([]string{"flash_sale", strconv.Itoa(unix)}, "-"), 0, 1000).Result()
	for _, slice := range slices {
		var product ProductCreate
		json.Unmarshal([]byte(slice), &product)
		if product.Id != "" {
			products = append(products, product)
		}
	}

	//json.Unmarshal([]byte(stringResult), &products)
	return products, nil
}

func (s *ProductModel) GetProductDetailMongo(id string) (ProductModel, error) {
	db := _mongo.ConnectMongo()
	c := db.Collection(collection)

	var productModel ProductModel
	condition := bson.M{"id": id}

	ctx := common.GetContext()
	err := c.FindOne(ctx, condition).Decode(&productModel)

	return productModel, err
}

func (s *ProductModel) SearchProductElastic(p ProductSearch) []interface{} {
	query := elastic.NewPrefixQuery("name", p.Name)
	searchResult := _elastic.Search(eIndex, query)

	var product ProductCreate
	var products []interface{}
	for _, item := range searchResult.Each(reflect.TypeOf(product)) {
		if t, ok := item.(ProductCreate); ok {
			products = append(products, t)
		}
	}

	return products
}
