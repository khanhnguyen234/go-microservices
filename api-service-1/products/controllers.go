package products

import (
	"github.com/gin-gonic/gin"
	"khanhnguyen234/api-service-1/common"
	"fmt"
	"net/http"
)

func ProductCreate (c *gin.Context) {
	var body ProductCreateRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := ProductModel{Name: body.Name, Price: body.Price}

	db := common.GetPostgreSQL()
	db.Create(&product)

	c.JSON(200, gin.H{"result": product})
}

func ProductList (c *gin.Context) {
	var products []ProductModel

	db := common.GetPostgreSQL()
	db.Find(&products)

	c.JSON(200, gin.H{"result": products})
}

func ProductDetail (c *gin.Context) {
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
	var query ProductFilterRequest

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	fmt.Println(query);

	var products []ProductModel
	db := common.GetPostgreSQL()
	db.Where("Price <= ?", query.Price).Find(&products)

	c.JSON(200, gin.H{"result": products})
}