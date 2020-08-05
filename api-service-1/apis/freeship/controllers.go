package freeship

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"khanhnguyen234/api-service-1/_rabbitmq"
)

func Publish(c *gin.Context) {
	body := FreeshipCreateRequest{
		ProductId:  "abc",
		IsFreeShip: true,
	}

	jsonBody, _ := json.Marshal(body)
	msg := string(jsonBody)

	_rabbitmq.LogsPublisher(msg)
	c.JSON(200, gin.H{"result": body})
}
