package products

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"khanhnguyen234/api-service-1/_postgres"
	"khanhnguyen234/api-service-1/_redis"
	"khanhnguyen234/api-service-1/apis/redis"
	"khanhnguyen234/api-service-1/utils"
	"net/http"
)

func ProductCache(c *gin.Context) {
	redis.Increase()
	var query ProductFilterRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	stringResult, err := _redis.Get(query.Price)
	if err != true {
		result := utils.ParseJsonToStruct(stringResult)
		c.JSON(200, gin.H{"isCache": true, "count": result["Count"]})
		return
	}

	var products []ProductModel
	var count int
	db := _postgres.GetPostgres()
	db.Where("Price <= ?", query.Price).Find(&products).Count(&count)

	result := make(map[string]interface{})
	result["Count"] = count

	redisValue := utils.ParseStructToJson(result)
	_redis.Set(query.Price, redisValue)
	c.JSON(200, gin.H{"isCache": false, "count": result["Count"]})
}

func ProductCreate(c *gin.Context) {
	redis.Increase()
	var body ProductCreateRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := ProductModel{Name: body.Name, Price: body.Price}

	db := _postgres.GetPostgres()

	db.Create(&product)
	elasticId := ElasticCreateProduct(product)
	PubProductCreated(product)

	c.JSON(200, gin.H{"result": product, "elasticId": elasticId})
}

func ProductList(c *gin.Context) {
	redis.Increase()
	var products []ProductModel

	db := _postgres.GetPostgres()
	db.Find(&products)

	c.JSON(200, gin.H{"result": products})
}

func ProductDetail(c *gin.Context) {
	redis.Increase()
	var uri ProductDetailRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	fmt.Println(uri)

	var product ProductModel
	db := _postgres.GetPostgres()
	db.Where("id = ?", uri.Id).First(&product)

	c.JSON(200, gin.H{"result": product})
}

func ProductFilter(c *gin.Context) {
	redis.Increase()
	var query ProductFilterRequest

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	fmt.Println(query)

	var products []ProductModel
	var count int
	db := _postgres.GetPostgres()
	db.Limit(10).Where("Price <= ?", query.Price).Find(&products).Count(&count)

	c.JSON(200, gin.H{"count": count, "result": products})
}

func ProductSearch(c *gin.Context) {
	redis.Increase()
	var uri ProductSearchRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	fmt.Println(uri)

	products := ElasticGetProductByName(uri.Name)

	c.JSON(200, gin.H{"result": products})
}

func DummyProduct() {
	db := _postgres.GetPostgres()
	for i := 0; i < 100000; i++ {
		price := utils.RandomInt(1000000)
		name := utils.RandomString(utils.RandomInt(1000))
		product := ProductModel{Name: name, Price: price}
		db.Create(&product)
		ElasticCreateProduct(product)
	}
	return
}
