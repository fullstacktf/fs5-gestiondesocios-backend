package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manolors/gorm-init-example/src/api/models"
	"github.com/manolors/gorm-init-example/src/api/utils"
)

//GetGame retrieves a game by id
func GetGame(writer http.ResponseWriter, r *http.Request) {
	game := models.Game{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&game, id)
	if game.ID > 0 {
		jGame, _ := json.Marshal(game)
		utils.SendResponse(writer, http.StatusOK, jGame)
	} else {
		utils.SendError(writer, http.StatusNotFound)
	}

}
