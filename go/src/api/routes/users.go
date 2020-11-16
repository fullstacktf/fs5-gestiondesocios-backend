package routes

import (
	"fs5-gestiondesocios-backend/src/api/controllers"

	"github.com/gorilla/mux"
)

//SetUsersRoutes sets the user routes
func SetUsersRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/users/{id}",
		controllers.GetUser).Methods("GET")
	subRouter.HandleFunc("/users",
		controllers.GetUsers).Methods("GET")
}
