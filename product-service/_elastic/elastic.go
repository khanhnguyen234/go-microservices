package _elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"khanhnguyen234/product-service/common"
	"log"
	"os"
)

var elasticClient *elastic.Client

func ConnectElastic() *elastic.Client {
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

func Put(index string, eType string, id string, data interface{}) error {
	ctx := context.Background()
	elasticClient := GetElastic()

	exists, err := elasticClient.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := elasticClient.CreateIndex(index).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
		}
	}

	_, err = elasticClient.Index().
		Index(index).
		Type(eType).
		Id(id).
		BodyJson(data).
		Do(ctx)

	if err != nil {
		log.Println(index, eType, id, data)
	}

	return err
}

func Search(index string, termQuery elastic.Query) *elastic.SearchResult {
	ctx := context.Background()
	elasticClient := GetElastic()

	searchResult, err := elasticClient.Search().
		Index(index).
		Query(termQuery).
		Pretty(true).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	return searchResult
}

func SearchBuilder(index string, q string) *elastic.SearchResult {
	ctx := context.Background()
	elasticClient := GetElastic()

	searchResult, err := elasticClient.Search().
		Index(index).
		Source(q).
		Pretty(true).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	return searchResult
}
