package tests

import (
	"fs5-gestiondesocios-backend/api/controllers"
	"fs5-gestiondesocios-backend/api/utils"
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

// games.go Tests
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
		Body(`{"id":1, "game_name":"juego1", "game_image":"https://cf.geekdo-images.com/NaVK216SnjDz3VLr5kKOAg__original/img/zcBNkPVorAebzEyRQNXr4dZkVsk=/0x0/pic350302.jpg" "rating":2, "id_owner":1,"entryDate":"10-10-2020","disponibility":false,"comments":"10/10"}`).
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

// End of games.go Tests

// assoc_partners.go Tests

func TestGetPartnersWorking(t *testing.T) {
	clearTable()
	insertGame()
	apitest.New().
		Debug().
		Handler(newApp().Router).
		Get("/api/assoc_partners").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetPartnerWorking(t *testing.T) {
	clearTable()
	insertGame()
	var getUserMock = apitest.NewMock().
		Get("/api/assoc_partners/1").
		RespondWith().
		Body(`{"id":1,"partner_name":"Pepe"}`).
		Status(http.StatusOK).
		End()

	apitest.New().
		Mocks(getUserMock).
		Handler(newApp().Router).
		Get("/api/assoc_partners/1").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetPartnerNotExist(t *testing.T) {
	clearTable()
	var getUserMock = apitest.NewMock().
		Get("/api/assoc_partners/1010").
		RespondWith().
		Body(`Partner not found`).
		Status(http.StatusNotFound).
		End()

	apitest.New().
		Mocks(getUserMock).
		Handler(newApp().Router).
		Get("/api/assoc_partners/1010").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

// End of assoc_partners.go Tests

// borrowed_games.go Tests

func TestGetBorrowedGamesWorking(t *testing.T) {
	clearTable()
	insertBorrowedGame()

	apitest.New().
		Debug().
		Handler(newApp().Router).
		Get("/api/borrowedgames").
		Expect(t).
		Status(http.StatusOK).
		End()
}

// End of borrowed_games.go Tests

type app struct {
	Router *mux.Router
}

func newApp() *app {
	router := mux.NewRouter()
	controllers.SetDatabase(utils.GetConnectionTest())
	subRouter := router.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/games/{id}",
		controllers.GetGame).Methods("GET")
	subRouter.HandleFunc("/games",
		controllers.GetGames).Methods("GET")
	subRouter.HandleFunc("/games",
		controllers.InsertGame).Methods("POST")
	subRouter.HandleFunc("/assoc_partners",
		controllers.GetPartners).Methods("GET")
	subRouter.HandleFunc("/assoc_partners/{id}",
		controllers.GetPartner).Methods("GET")
	subRouter.HandleFunc("/assoc_partners",
		controllers.InsertPartner).Methods("POST")
	subRouter.HandleFunc("/borrowedgames",
		controllers.GetBorrowedGames).Methods("GET")
	return &app{Router: router}
}
