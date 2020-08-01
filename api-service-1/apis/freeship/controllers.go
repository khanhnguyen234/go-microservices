package freeship

import (
	"github.com/gin-gonic/gin"
	"khanhnguyen234/api-service-1/rabbitmq"
	"encoding/json"
)

func Publish (c *gin.Context) {
	body := FreeshipCreateRequest{
		ProductId: "abc",
		IsFreeShip: true,
	}

	jsonBody, _ := json.Marshal(body)
	msg := string(jsonBody)

	rabbitmq.LogsPublisher(msg)
	c.JSON(200, gin.H{"result": body})
}

