package _elastic

import (
	"os"
	"github.com/olivere/elastic/v7"
	"khanhnguyen234/api-service-1/common"
)

var elasticClient *elastic.Client

func ConnectElastic()  *elastic.Client {
	var err error
	ELASTIC_URL := os.Getenv("ELASTIC_URL")

	elasticClient, err = elastic.NewClient(
		elastic.SetURL(ELASTIC_URL),
		elastic.SetSniff(false),
	)
	common.LogStatus(err, "Connect Elastic")
	return elasticClient
}

func GetElastic() *elastic.Client {
	return elasticClient
}
