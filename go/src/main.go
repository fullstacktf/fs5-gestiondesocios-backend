package main

import (
	"fs5-gestiondesocios-backend/src/api/routes"
	"fs5-gestiondesocios-backend/src/api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	utils.MigrateDB()
	router := mux.NewRouter()
	routes.SetUsersRoutes(router)
	routes.SetGamesRoutes(router)
	routes.SetBorrowedGamesRoutes(router)
	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}
	log.Println("Running on port 80")
	log.Println(server.ListenAndServe())
}

//To generate the docs use this command /Users/daviddiaz/go/bin/golds -gen -emphasize-wdpkgs -compact .
