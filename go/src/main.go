package main

import (
	"fs5-gestiondesocios-backend/api/routes"
	"fs5-gestiondesocios-backend/api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                            // All origins
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"}, // Allowing only get, just an example
	})
	utils.MigrateDB()
	router := mux.NewRouter()
	routes.SetUsersRoutes(router)
	routes.SetGamesRoutes(router)
	routes.SetBorrowedGamesRoutes(router)
	routes.SetAssocPartnerRoutes(router)
	server := http.Server{
		Addr:    ":80",
		Handler: c.Handler(router),
	}
	utils.Populate_db()
	log.Println("Running on port 80")
	log.Println(server.ListenAndServe())
}

//To generate the docs use this command /Users/daviddiaz/go/bin/golds -gen -emphasize-wdpkgs -compact .
