package shop

import (
	"github.com/khanhnguyen234/go-microservices/_common"
	"github.com/khanhnguyen234/go-microservices/_elastic"
	"github.com/khanhnguyen234/go-microservices/_mongo"
	"github.com/olivere/elastic"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"time"
)

const (
	collection = "shop"
	eIndex     = "shop"
	eType      = "byName"
)

func (s *ShopCreate) CreateShopMongo(shop ShopCreate) (ShopCreate, error) {
	if shop.Id == "" {
		shop.Id = uuid.NewV4().String()
	}
	shop.CreatedAt = time.Now()

	err := _mongo.InsertOne(collection, shop)

	return shop, err
}

func (s *ShopCreate) CreateShopElastic(shop ShopCreate) error {
	err := _elastic.Put(eIndex, eType, shop.Id, shop)
	return err
}

func (s *ShopModel) GetShopsMongo() ([]ShopModel, error) {
	db := _mongo.ConnectMongo()
	c := db.Collection(collection)

	ctx := _common.GetContext()
	condition := bson.D{}
	cur, err := c.Find(ctx, condition)

	var shops []ShopModel

	for cur.Next(ctx) {
		var t ShopModel
		err := cur.Decode(&t)
		if err != nil {
			return shops, err
		}

		shops = append(shops, t)
	}

	if err := cur.Err(); err != nil {
		return shops, err
	}

	return shops, err
}

func (s *ShopModel) GetShopDetailMongo(id string) (ShopModel, error) {
	db := _mongo.ConnectMongo()
	c := db.Collection(collection)

	var shopModel ShopModel
	condition := bson.M{"id": id}

	ctx := _common.GetContext()
	err := c.FindOne(ctx, condition).Decode(&shopModel)

	return shopModel, err
}

func (s *ShopModel) SearchShopElastic(p ShopSearch) []interface{} {
	query := elastic.NewPrefixQuery("name", p.Name)
	searchResult := _elastic.Search(eIndex, query)

	//search := fmt.Sprintf(`
	//{
	//	"query": {
	//		"prefix": {
	//			"name": {
	//				"value": "%s"
	//			}
	//		}
	//	}
	//}
	//`, p.Name)
	//searchResult := _elastic.SearchBuilder(eIndex, search)

	var shop ShopCreate
	var shops []interface{}
	for _, item := range searchResult.Each(reflect.TypeOf(shop)) {
		if t, ok := item.(ShopCreate); ok {
			shops = append(shops, t)
		}
	}

	return shops
}
