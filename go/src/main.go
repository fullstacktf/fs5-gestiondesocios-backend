package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func main() {
	config := &DBConfig{
		Host:     "app_mariadb",
		Port:     3306,
		User:     "dev",
		DBName:   "fullstackAsociacion",
		Password: "passdev",
	}
	_, err := gorm.Open(mysql.Open(DbURL(config)), &gorm.Config{})
	if err != nil {
		log.Fatal("Error BBDD")
	} else {
		log.Printf("It Worked!")
	}
}
