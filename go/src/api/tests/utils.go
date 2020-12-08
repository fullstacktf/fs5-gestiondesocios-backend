package tests

import (
	"fs5-gestiondesocios-backend/src/api/utils"
	"log"

	"github.com/gorilla/mux"
)

var db = utils.GetConnectionTest()
var router = mux.NewRouter()

func assertTableExists() {
	sql, _ := db.DB()
	if _, err := sql.Exec(tableCreationQueryAssocPartners); err != nil {
		log.Fatal(err)
	}
	if _, err := sql.Exec(tableCreationQueryGames); err != nil {
		log.Fatal(err)
	}
	if _, err := sql.Exec(tableCreationQueryBorrowedGames); err != nil {
		log.Fatal(err)
	}
	if _, err := sql.Exec(tableCreationQueryAssocUsers); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQueryAssocPartners = `CREATE TABLE IF NOT EXISTS assoc_partners (
	id INT PRIMARY KEY,
	partner_name VARCHAR(30) NOT NULL
  );`

const tableCreationQueryGames = `CREATE TABLE IF NOT EXISTS games (
	id INT PRIMARY KEY,
	game_name VARCHAR(200),
	image VARCHAR(200),
  	rating  FLOAT,
	id_owner INT,
	entry_date VARCHAR(200),
	disponibility BOOL NOT NULL,
	comments VARCHAR(200),
	CONSTRAINT fk_idOwner FOREIGN KEY (id_owner) REFERENCES assoc_partners (id)
  );`

const tableCreationQueryBorrowedGames = `CREATE TABLE IF NOT EXISTS borrowed_games (
	id_game INT PRIMARY KEY,
	id_borrower INT NOT NULL,
	borrow_date VARCHAR(200),
	FOREIGN KEY (id_borrower) REFERENCES assoc_partners (id),
	CONSTRAINT fk_idGame FOREIGN KEY (id_game) REFERENCES games (id)
  );`

const tableCreationQueryAssocUsers = `CREATE TABLE IF NOT EXISTS assoc_users (
	id INT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	user_password VARCHAR(30) NOT NULL
  );`

const insertGameQuery = `INSERT INTO games VALUES (1, "juego1", "https://cf.geekdo-images.com/NaVK216SnjDz3VLr5kKOAg__original/img/zcBNkPVorAebzEyRQNXr4dZkVsk=/0x0/pic350302.jpg", 2, 1, "10-10-2020", true, "10/10");`
const insertAssocPartnerQuery = `INSERT INTO assoc_partners VALUES (1, "Pepe");`

func clearTable() {
	db.Exec("DELETE FROM borrowed_games")
	db.Exec("DELETE FROM games")
	db.Exec("DELETE FROM assoc_partners")
	db.Exec("DELETE FROM assoc_users")
}

func insertGame() {
	db.Exec(insertAssocPartnerQuery)
	db.Exec(insertGameQuery)
}

const insertUnavailableGameQuery = `INSERT INTO games VALUES (1, "juego1", "https://cf.geekdo-images.com/NaVK216SnjDz3VLr5kKOAg__original/img/zcBNkPVorAebzEyRQNXr4dZkVsk=/0x0/pic350302.jpg", 2, 1, "10-10-2020", false, "10/10");`
const insertBorrowedGameQuery = `INSERT INTO borrowed_games VALUES (1, 1, "10-10-2020");`

func insertBorrowedGame() {
	db.Exec(insertAssocPartnerQuery)
	db.Exec(insertUnavailableGameQuery)
	db.Exec(insertBorrowedGameQuery)
}
