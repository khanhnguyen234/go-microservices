package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"khanhnguyen234/product-service/_mongo"
	"khanhnguyen234/product-service/apis/product"
	"khanhnguyen234/product-service/common"
	"log"
	"os"
)

const (
	collection = "product"
)

func ProductExport() {
	err := godotenv.Load()
	db := _mongo.ConnectMongo()
	c := db.Collection(collection)

	ctx := common.GetContext()
	condition := bson.D{}

	options := options.Find()
	options.SetSort(bson.D{{"updated_at", -1}})

	cur, err := c.Find(ctx, condition, options)
	if err != nil {
		panic(err)
	}

	var products []product.ProductCreate

	for cur.Next(ctx) {
		var t product.ProductCreate
		err := cur.Decode(&t)
		if err != nil {
			panic(err)
		}

		products = append(products, t)
	}

	file, _ := json.MarshalIndent(products, "", " ")

	_ = ioutil.WriteFile("backup/product.json", file, 0644)
}

func ProductImport() {
	jsonFile, err := os.Open("backup/product.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var products []interface{}
	json.Unmarshal(byteValue, &products)

	err = godotenv.Load()
	db := _mongo.ConnectMongo()
	c := db.Collection(collection)
	ctx := common.GetContext()
	_, err = c.InsertMany(ctx, products)

	defer jsonFile.Close()
}

func main() {
	argFunc := os.Args[1]

	switch argFunc {
	case "ProductImport":
		log.Println("Start ProductImport")
		ProductImport()
		log.Println("End ProductImport")
	case "ProductExport":
		log.Println("Start ProductExport")
		ProductExport()
		log.Println("End ProductExport")
	default:
		fmt.Println("Nothing Execute")
	}
}
