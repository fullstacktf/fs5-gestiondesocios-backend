package controllers

import (
	"encoding/json"
	"fs5-gestiondesocios-backend/src/api/models"
	"fs5-gestiondesocios-backend/src/api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetUser retrieves an user by id
func GetUser(writer http.ResponseWriter, r *http.Request) {
	user := models.AssocUser{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&user, id)
	if user.ID > 0 {
		jUser, _ := json.Marshal(user)
		utils.SendResponse(writer, http.StatusOK, jUser)
	} else {
		utils.SendError(writer, http.StatusNotFound)
	}

}

//GetUsers retrieves all users
func GetUsers(writer http.ResponseWriter, r *http.Request) {
	users := []models.AssocUser{}

	db := utils.GetConnection()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&users)
	jUsers, _ := json.Marshal(users)
	utils.SendResponse(writer, http.StatusOK, jUsers)
}
