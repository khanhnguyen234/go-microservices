package free_ship

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	body := FreeshipCreateRequest{
		ProductId:  "abc",
		IsFreeShip: true,
	}

	jsonBody, _ := json.Marshal(body)
	msg := string(jsonBody)
	PubFreeShip(msg)

	c.JSON(200, gin.H{"result": body})
}
