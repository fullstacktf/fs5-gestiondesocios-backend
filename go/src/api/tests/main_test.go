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
	apitest.New().
		Debug().
		Handler(newApp().Router).
		Get("/api/games").
		Expect(t).
		Status(http.StatusOK).
		End()
}

type app struct {
	Router *mux.Router
}

func newApp() *app {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/games",
		controllers.GetGames).Methods("GET")
	return &app{Router: router}
}
