package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"khanhnguyen234/api-gateway/_mongo"
	"khanhnguyen234/api-gateway/apis/auth"
	"khanhnguyen234/api-gateway/common"
	"net/http"
)

func main() {
	initRouter()
}

func initRouter() {
	err := godotenv.Load()
	common.LogStatus(err, "Load Env")

	_mongo.ConnectMongo()

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	basePath := route.Group("/")
	auth.AuthRouters(basePath.Group("/auth"))

	route.GET("/api_service_1", func(c *gin.Context) {
		response, err := http.Get("http://localhost:7001/query?name=query&id=7001")
		if err != nil {
			c.JSON(400, gin.H{"err": err})
		} else {
			var person PersonParam
			data, _ := ioutil.ReadAll(response.Body)
			stringJson := string(data)
			json.Unmarshal([]byte(stringJson), &person)

			fmt.Println(stringJson)
			fmt.Println(person)

			c.JSON(200, gin.H{"data": person})
		}
	})

	route.GET("/api_service_1/products/filter", func(c *gin.Context) {
		var query ProductFilterQuery

		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		response, err := http.Get("http://localhost:7001/no-auth/products/filter?price=" + query.Price)

		if err != nil {
			c.JSON(400, gin.H{"err": err})
		} else {
			var result map[string]interface{}
			data, _ := ioutil.ReadAll(response.Body)
			stringJson := string(data)
			json.Unmarshal([]byte(stringJson), &result)

			c.JSON(200, gin.H{"result": result["result"]})
		}
	})

	common.LogSuccess("Listening and serving HTTP on :7000")
	route.Run(":7000")
}

type PersonParam struct {
	Name string
	Id   string
}

type ProductFilterQuery struct {
	Name  string `form:"name"`
	Price string `form:"price"`
}
