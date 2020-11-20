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
    idOwner INT,
    entryDate DATE,
    disponibility BOOL NOT NULL,
    comments VARCHAR(200),
    CONSTRAINT fk_idOwner FOREIGN KEY (idOwner) REFERENCES assoc_partners (id)
  );`

func clearTable() {
	db.Exec("DELETE FROM assoc_partners")
}
