package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"khanhnguyen234/job/_mongo"
	"khanhnguyen234/job/_rabbitmq"
	"khanhnguyen234/job/common"
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
	common.LogStatus(err, "Load Env")

	_, err = _rabbitmq.ConnectRabbitMQ()
	_, err = _mongo.ConnectMongo()

	if err != nil {
		return
	}

	eventTrigger()
	scheduleTrigger()

	<-gocron.Start()
}
