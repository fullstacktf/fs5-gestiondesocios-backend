package controllers

import "gorm.io/gorm"

var db *gorm.DB
var initDB bool

// SetDatabase selects database for testing
func SetDatabase(database *gorm.DB) {
	initDB = true
	db = database
}
