package utils

import (
	"fmt"
	"log"

	"github.com/manolors/gorm-init-example/src/api/models"
)

//MigrateDB automigrates the DB from the model
func MigrateDB() {
	db := GetConnection()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	fmt.Println("Migrating models...")
	db.AutoMigrate(&models.Games{})
}
