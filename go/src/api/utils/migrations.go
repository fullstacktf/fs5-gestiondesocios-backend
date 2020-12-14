package utils

import (
	"fmt"
	"fs5-gestiondesocios-backend/api/models"
	"log"
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
	db.AutoMigrate(&models.Game{})
	db.AutoMigrate(&models.AssocUser{})
}
