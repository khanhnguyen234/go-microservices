package _postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"khanhnguyen234/api-service-1/common"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func ConnectPostgres() *gorm.DB {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", 
	os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"))
	
	db, err := gorm.Open("postgres", conn)
	common.LogStatus(err, "Connect Postgres")
	
	db.DB().SetMaxIdleConns(10)
	// db.LogMode(true)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetPostgres() *gorm.DB {
	return DB
}

// This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("postgres", "./../gorm_test.db")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

// Delete the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("./../gorm_test.db")
	return err
}