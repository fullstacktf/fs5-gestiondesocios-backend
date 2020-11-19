package main

import (
	"fmt"
	"fs5-gestiondesocios-backend/src/api/routes"
	"fs5-gestiondesocios-backend/src/api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type dBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func dbURL(dbConfig *dBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func main() {
	utils.MigrateDB()
	router := mux.NewRouter()
	routes.SetUsersRoutes(router)
	routes.SetGamesRoutes(router)
	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}
	log.Println("Running on port 80")
	log.Println(server.ListenAndServe())
	/*atomic.AddUint64(&counter, 0)
	router := mux.NewRouter()
	router.HandleFunc("/", Counter).Methods("GET")
	fmt.Println("GO REST server running on http://localhost:8080 ")
	log.Fatal(http.ListenAndServe(":80", router))*/
}

//To generate the docs use this command /Users/daviddiaz/go/bin/golds -gen -emphasize-wdpkgs -compact .
