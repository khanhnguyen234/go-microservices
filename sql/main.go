package main

import (
	"github.com/joho/godotenv"
	"khanhnguyen234/sql/_postgres"
	"khanhnguyen234/sql/common"
	"khanhnguyen234/sql/function"
	"khanhnguyen234/sql/procedure"
	"khanhnguyen234/sql/trigger"
	"log"
)

func main() {
	err := godotenv.Load()
	common.LogStatus(err, "Load Env")
	_postgres.ConnectPostgres()

	function.ProductSearch()
	function.ExecProductSearch("pro")

	procedure.ProductSearch()
	procedure.ExecProductSearch("pro")

	trigger.ProductCreate()

	forever := make(chan bool)

	log.Printf("To exit press CTRL+C")

	<-forever

}
