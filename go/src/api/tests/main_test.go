package tests

import (
	"fs5-gestiondesocios-backend/src/api/controllers"
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
)

func TestMain(m *testing.M) {
	assertTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestGetGamesWorking(t *testing.T) {
	clearTable()
	insertGame()
	apitest.New().
		Debug().
		Handler(newApp().Router).
		Get("/api/games").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetGameWorking(t *testing.T) {
	clearTable()
	insertGame()
	var getUserMock = apitest.NewMock().
		Get("/api/games/1").
		RespondWith().
		Body(`{"id":1,"idOwner":1,"entryDate":"10-10-2020","disponibility":false,"comments":"10/10"}`).
		Status(http.StatusOK).
		End()

	apitest.New().
		Mocks(getUserMock).
		Handler(newApp().Router).
		Get("/api/games/1").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetGameNotExist(t *testing.T) {
	clearTable()
	var getUserMock = apitest.NewMock().
		Get("/api/games/1010").
		RespondWith().
		Body(`Game not found`).
		Status(http.StatusNotFound).
		End()

	apitest.New().
		Mocks(getUserMock).
		Handler(newApp().Router).
		Get("/api/games/1010").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

type app struct {
	Router *mux.Router
}

func newApp() *app {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/games/{id}",
		controllers.GetGame).Methods("GET")
	subRouter.HandleFunc("/games",
		controllers.GetGames).Methods("GET")
	subRouter.HandleFunc("/games",
		controllers.InsertGame).Methods("POST")
	return &app{Router: router}
}
