package products

import (
	"context"
	"fmt"
	"khanhnguyen234/api-service-1/_elastic"
	"reflect"
)

const (
	elasticIndexName = "product"
	elasticTypeName  = "byName"
)

func ElasticCreateProduct(product ProductModel) string {
	ctx := context.Background()
	elasticClient := _elastic.GetElastic()

	exists, err := elasticClient.IndexExists(elasticIndexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := elasticClient.CreateIndex(elasticIndexName).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
		}
	}

	put, err := elasticClient.Index().
		Index(elasticIndexName).
		Type(elasticTypeName).
		Id(product.Name).
		BodyJson(product).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	return put.Id
}

func ElasticGetProductByName(name string) []ProductModel {
	ctx := context.Background()
	elasticClient := _elastic.GetElastic()

	_search := fmt.Sprintf(`
	{
		"query": {
			"prefix": {
				"Name": {
					"value": "%s"
				}
			}
		}
	}
	`, name)
	fmt.Println(_search)

	searchResult, err := elasticClient.Search().
		Index(elasticIndexName).
		Source(_search).
		Pretty(true).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	var product ProductModel
	var products []ProductModel

	for _, item := range searchResult.Each(reflect.TypeOf(product)) {
		if t, ok := item.(ProductModel); ok {
			products = append(products, t)
		}
	}

	return products
}
