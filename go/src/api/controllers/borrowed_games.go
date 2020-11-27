package controllers

import (
	"encoding/json"
	"fmt"
	"fs5-gestiondesocios-backend/src/api/models"
	"fs5-gestiondesocios-backend/src/api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetBorrowedGames retrieves all borrowed games by idGame
func GetBorrowedGames(writer http.ResponseWriter, r *http.Request) {
	borrowedGames := []models.BorrowedGame{}

	db := utils.GetConnection()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&borrowedGames)
	if len(borrowedGames) > 0 {
		jBorrowedGames, _ := json.Marshal(borrowedGames)
		utils.SendResponse(writer, http.StatusOK, jBorrowedGames)
	} else {
		utils.SendError(writer, http.StatusNotFound, "Borrowed Games not found")
	}

}

// InsertBorrowedGame inserts a borrowed game in the table
func InsertBorrowedGame(writer http.ResponseWriter, r *http.Request) {
	borrowedGame := models.BorrowedGame{}
	db := utils.GetConnection()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	errDB := json.NewDecoder(r.Body).Decode(&borrowedGame)
	if errDB != nil {
		fmt.Println(errDB)
		utils.SendError(writer, http.StatusBadRequest, "Error formating game")
		return
	}

	errDB = db.Create(&borrowedGame).Error
	if errDB != nil {
		fmt.Println(errDB)
		utils.SendError(writer, http.StatusBadRequest, "Error inserting borrowedGame")
		return
	}

	gameToBorrow := models.Game{}
	db.Find(&gameToBorrow, borrowedGame.IDGame)
	db.Model(&gameToBorrow).UpdateColumn("disponibility", false)

	j, _ := json.Marshal(borrowedGame)
	utils.SendResponse(writer, http.StatusCreated, j)
}

// DeleteBorrowedGame deletes a game from the table and updates disponibility
func DeleteBorrowedGame(writer http.ResponseWriter, r *http.Request) {
	borrowedGame := models.BorrowedGame{}
	id := mux.Vars(r)["id_game"]
	db := utils.GetConnection()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&borrowedGame, id)
	if borrowedGame.IDGame > 0 {
		db.Delete(borrowedGame)
		utils.SendResponse(writer, http.StatusOK, []byte(`Borrowed game removed successfully`))
		gameToDelete := models.Game{}
		db.Find(&gameToDelete, borrowedGame.IDGame)
		db.Model(&gameToDelete).UpdateColumn("disponibility", true)
	} else {
		utils.SendError(writer, http.StatusNotFound, "Error removing the borrowed game, might not exist")
	}

}
