package products

import (
	"github.com/gin-gonic/gin"
	"khanhnguyen234/api-service-1/common"
	"khanhnguyen234/api-service-1/utils"
	"khanhnguyen234/api-service-1/apis/redis"
	"fmt"
	"net/http"
)

func ProductCache (c *gin.Context) {
	redis.Increase()
	var query ProductFilterRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	stringResult, err := common.GetRedis(query.Price)
	if (err != true) {
		result := utils.ParseJsonToStruct(stringResult);
		c.JSON(200, gin.H{"isCache": true, "count": result["Count"]})
		return
	}

	var products []ProductModel
	var count int
	db := common.GetPostgreSQL()
	db.Where("Price <= ?", query.Price).Find(&products).Count(&count)

	result := make(map[string]interface{})
	result["Count"] = count

	redisValue := utils.ParseStructToJson(result)
	common.SetRedis(query.Price, redisValue)
	c.JSON(200, gin.H{"isCache": false, "count": result["Count"]})
}

func ProductCreate (c *gin.Context) {
	redis.Increase()
	var body ProductCreateRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := ProductModel{Name: body.Name, Price: body.Price}

	db := common.GetPostgreSQL()
	db.Create(&product)
	elasticId := ElasticCreateProduct(product)

	c.JSON(200, gin.H{"result": product, "elasticId": elasticId})
}

func ProductList (c *gin.Context) {
	redis.Increase()
	var products []ProductModel

	db := common.GetPostgreSQL()
	db.Find(&products)

	c.JSON(200, gin.H{"result": products})
}

func ProductDetail (c *gin.Context) {
	redis.Increase()
	var uri ProductDetailRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	fmt.Println(uri);

	var product ProductModel
	db := common.GetPostgreSQL()
	db.Where("id = ?", uri.Id).First(&product)

	c.JSON(200, gin.H{"result": product})
}

func ProductFilter (c *gin.Context) {
	redis.Increase()
	var query ProductFilterRequest

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	fmt.Println(query);

	var products []ProductModel
	var count int
	db := common.GetPostgreSQL()
	db.Limit(10).Where("Price <= ?", query.Price).Find(&products).Count(&count)

	c.JSON(200, gin.H{"count": count, "result": products})
}

func ProductSearch (c *gin.Context) {
	redis.Increase()
	var uri ProductSearchRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	fmt.Println(uri);

	products := ElasticGetProductByName(uri.Name)


	c.JSON(200, gin.H{"result": products})
}

func DummyProduct() {
	db := common.GetPostgreSQL()
	for i := 0; i < 100000; i++ {
		price := utils.RandomInt(1000000)
		name := utils.RandomString(utils.RandomInt(1000))
		product := ProductModel{Name: name, Price: price}
		db.Create(&product)
		ElasticCreateProduct(product)
	}
	return
}