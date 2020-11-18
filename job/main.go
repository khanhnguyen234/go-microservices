package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/khanhnguyen234/go-microservices/_common"
	"github.com/khanhnguyen234/go-microservices/_mongo"
	"github.com/khanhnguyen234/go-microservices/_rabbitmq"
	"khanhnguyen234/job/services/email"
)

func eventTrigger() {
	email.SubProductCreate()
}

func scheduleTrigger() {
	email.SendEmail()
}

func main() {
	err := godotenv.Load()
	_common.LogStatus(err, "Load Env")

	_, err = _rabbitmq.ConnectRabbitMQ()
	_mongo.ConnectMongo()

	if err != nil {
		return
	}

	eventTrigger()
	scheduleTrigger()

	<-gocron.Start()
}
