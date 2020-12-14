package controllers

import (
	"encoding/json"
	"fmt"
	"fs5-gestiondesocios-backend/api/models"
	"fs5-gestiondesocios-backend/api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetPartner retrieves a single partner from table
func GetPartner(writer http.ResponseWriter, r *http.Request) {
	partner := models.AssocPartner{}
	id := mux.Vars(r)["id"]

	if !initDB {
		db = utils.GetConnection()
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&partner, id)
	if partner.ID > 0 {
		jpartner, _ := json.Marshal(partner)
		utils.SendResponse(writer, http.StatusOK, jpartner)
	} else {
		utils.SendError(writer, http.StatusNotFound, "Partner not found")
	}

}

//GetPartners retrieves all partners by id
func GetPartners(writer http.ResponseWriter, r *http.Request) {
	partners := []models.AssocPartner{}

	if !initDB {
		db = utils.GetConnection()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&partners)
	if len(partners) > 0 {
		jpartners, _ := json.Marshal(partners)
		utils.SendResponse(writer, http.StatusOK, jpartners)
	} else {
		utils.SendError(writer, http.StatusNotFound, "Partners not found")
	}

}

//InsertPartner inserts a partner into the "assoc_partners" table
func InsertPartner(writer http.ResponseWriter, r *http.Request) {
	partner := models.AssocPartner{}
	if !initDB {
		db = utils.GetConnection()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	errDB := json.NewDecoder(r.Body).Decode(&partner)
	if errDB != nil {
		fmt.Println(errDB)
		utils.SendError(writer, http.StatusBadRequest, "Error formating assoc_partners")
		return
	}

	errDB = db.Create(&partner).Error
	if errDB != nil {
		fmt.Println(errDB)
		utils.SendError(writer, http.StatusBadRequest, "Error inserting partner")
		return
	}

	jpartner, _ := json.Marshal(partner)
	utils.SendResponse(writer, http.StatusCreated, jpartner)
}

//DeletePartner deletes a partner from table
func DeletePartner(writer http.ResponseWriter, r *http.Request) {
	partner := models.AssocPartner{}
	id := mux.Vars(r)["id"]
	if !initDB {
		db = utils.GetConnection()
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&partner, id)
	if partner.ID > 0 {
		db.Delete(partner)
		utils.SendResponse(writer, http.StatusOK, []byte(`Partner removed successfully`))
	} else {
		utils.SendError(writer, http.StatusNotFound, "Error removing the partner, might not exist")
	}
}

//UpdatePartner updates an existing partner from the table
func UpdatePartner(writer http.ResponseWriter, r *http.Request) {
	partnerFind := models.AssocPartner{}
	partnerNewData := models.AssocPartner{}

	id := mux.Vars(r)["id"]
	if !initDB {
		db = utils.GetConnection()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error clossing the DB")
	} else {
		defer sqlDB.Close()
	}
	db.Find(&partnerFind, id)

	if partnerFind.ID > 0 {
		err := json.NewDecoder(r.Body).Decode(&partnerNewData)
		if err != nil {
			utils.SendError(writer, http.StatusBadRequest, "Invalid data")
			return
		}
		db.Model(&partnerFind).Updates(partnerNewData)
		j, _ := json.Marshal(partnerFind)
		utils.SendResponse(writer, http.StatusOK, j)
	} else {
		utils.SendError(writer, http.StatusNotFound, "Partner not found")
	}
}
