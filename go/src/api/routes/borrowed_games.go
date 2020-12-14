package routes

import (
	"fs5-gestiondesocios-backend/api/controllers"

	"github.com/gorilla/mux"
)

//SetBorrowedGamesRoutes sets the BorrowedGames routes
func SetBorrowedGamesRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/borrowedgames",
		controllers.GetBorrowedGames).Methods("GET")
	subRouter.HandleFunc("/borrowedgames",
		controllers.InsertBorrowedGame).Methods("POST")
	subRouter.HandleFunc("/borrowedgames/{id_game}",
		controllers.DeleteBorrowedGame).Methods("DELETE")
}
