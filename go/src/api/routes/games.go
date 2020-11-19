package routes

import (
	"fs5-gestiondesocios-backend/src/api/controllers"

	"github.com/gorilla/mux"
)

//SetGamesRoutes sets the game routes
func SetGamesRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/games/{id}",
		controllers.GetGame).Methods("GET")
	subRouter.HandleFunc("/games",
		controllers.InsertGame).Methods("POST")
}
