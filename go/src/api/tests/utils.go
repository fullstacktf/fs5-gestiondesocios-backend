package tests

import (
	"fs5-gestiondesocios-backend/src/api/utils"
	"log"

	"github.com/gorilla/mux"
)

var db = utils.GetConnection()
var router = mux.NewRouter()

func assertTableExists() {
	sql, _ := db.DB()
	if _, err := sql.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS games (
    id INT PRIMARY KEY,
    id_owner INT,
    entry_date VARCHAR(200),
    disponibility BOOL NOT NULL,
    comments VARCHAR(200),
    CONSTRAINT fk_idOwner FOREIGN KEY (idOwner) REFERENCES assoc_partners (id)
  );`

const insertGameQuery = `INSERT INTO games VALUES (1, 1, "10-10-2020", true, "10/10");`
const insertAssocPartnerQuery = `INSERT INTO assoc_partners VALUES (1, "Pepe");`

func clearTable() {
	db.Exec("DELETE FROM borrowed_games")
	db.Exec("DELETE FROM games")
	db.Exec("DELETE FROM assoc_partners")
}

func insertGame() {
	db.Exec(insertAssocPartnerQuery)
	db.Exec(insertGameQuery)
}

const insertUnavailableGameQuery = `INSERT INTO games VALUES (1, 1, "10-10-2020", false, "10/10");`
const insertBorrowedGameQuery = `INSERT INTO borrowed_games VALUES (1, 1, "10-10-2020");`

func insertBorrowedGame() {
	db.Exec(insertAssocPartnerQuery)
	db.Exec(insertUnavailableGameQuery)
	db.Exec(insertBorrowedGameQuery)
}
