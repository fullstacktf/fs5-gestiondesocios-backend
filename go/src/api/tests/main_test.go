package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	assertTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

//TestEmptyTable checks if doing a query on an empty table throws the correct error
func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/api/games", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

//
func TestGetNonExistentGame(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/api/games/11", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	if strings.EqualFold(response.Body.String(), "404 page not found") {
		t.Error("Response: ", response.Body.String(), "404 page not found")
		t.Errorf("Expected the 'error' key of the response to be set to '404 page not found'. Got '%s'", response.Body.String())
	}
}

func TestCreateGame(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{
        "id": 2,
        "idOwner": 1,
        "entryDate": "10-10-2020",
        "disponibiliry": true,
        "comments": "10/10"
    }`)
	req, _ := http.NewRequest("POST", "/api/games", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1 {
		t.Errorf("Expected game id to be '2'. Got '%v'", m["id"])
	}

	if m["idOwner"] != 1 {
		t.Errorf("Expected owner id to be '1'. Got '%v'", m["owner"])
	}

	if m["entryDate"] != "10-10-2020" {
		t.Errorf("Expected entry date to be '10-10-2020'. Got '%v'", m["entryDate"])
	}

	if m["disponibility"] != true {
		t.Errorf("Expected disponibility to be 'true'. Got '%v'", m["disponibility"])
	}
	if m["comments"] != true {
		t.Errorf("Expected comments to be '10/10'. Got '%v'", m["comments"])
	}

}
