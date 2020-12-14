package routes

import (
	"fs5-gestiondesocios-backend/api/controllers"

	"github.com/gorilla/mux"
)

//SetAssocPartnerRoutes sets the AssocPartners routes
func SetAssocPartnerRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/assoc_partners/{id}",
		controllers.GetPartner).Methods("GET")
	subRouter.HandleFunc("/assoc_partners",
		controllers.GetPartners).Methods("GET")
	subRouter.HandleFunc("/assoc_partners",
		controllers.InsertPartner).Methods("POST")
	subRouter.HandleFunc("/assoc_partners/{id}",
		controllers.DeletePartner).Methods("DELETE")
	subRouter.HandleFunc("/assoc_partners/{id}",
		controllers.UpdatePartner).Methods("PUT")
}
