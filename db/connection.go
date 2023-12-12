package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DNS = "host=localhost user=fazt password=mysecretpassword dbname=gorm port=5432"
var DB *gorm.DB

func DbConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in database connection: ", err)
	} else {
		log.Println("DB Connected ", DB)
	}
}
