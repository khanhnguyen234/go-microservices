package common

import (
	"fmt"
	"log"
	"time"
	"github.com/olivere/elastic/v7"
)

var elasticClient *elastic.Client

func InitElasticsearch()  *elastic.Client {
	var err error

	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://localhost:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			fmt.Println("Retry Elasticsearch....")
		} else {
			break
		}
	}

	return elasticClient
}

func Elasticsearch() *elastic.Client {
	return elasticClient
}
