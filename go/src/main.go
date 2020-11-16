package main

import (
	"encoding/json"
	"fmt"
	"fs5-gestiondesocios-backend/src/api/routes"
	"fs5-gestiondesocios-backend/src/api/utils"
	"log"
	"net/http"
	"sync/atomic"

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

var counter uint64

func Counter(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&counter, 1)
	json.NewEncoder(w).Encode(counter)
}
