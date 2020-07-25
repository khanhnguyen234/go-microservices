package main

import (
	"fmt"
	"khanhnguyen234/api-gateway/database/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func main() {
	initRouter()
}

func initRouter() {

	r := gin.Default()

	r.POST("/get_info_user", models.GetInfoUser)

	r.POST("/insert_user", models.InsertUser)

	r.GET("/api_service_1", func (c *gin.Context) {
		// response, err := http.Get("http://localhost:7001/param/param/7001")
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

	r.GET("/api_service_1/products/filter", func (c *gin.Context) {
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

	r.Run(":7000")
}

type PersonParam struct {
	Name string
	Id string
}

type ProductFilterQuery struct {
	Name string `form:"name"`
	Price string `form:"price"`
}