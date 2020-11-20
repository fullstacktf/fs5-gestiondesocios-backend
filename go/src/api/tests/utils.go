package tests

import (
	"fs5-gestiondesocios-backend/src/api/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var db = utils.GetConnectionLocalHost()
var router = mux.NewRouter()

func assertTableExists() {
	sql, _ := db.DB()
	if _, err := sql.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, but instead got %d\n", expected, actual)
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
