package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/graphql", graphql).Methods("GET", "POST", "OPTIONS", "PUT", "DELETE")
	return router
}

func disableCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-Type", "*")
}

func graphql(w http.ResponseWriter, r *http.Request){
	disableCors(&w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	disableCors(&w, r)

	if (*r).Method == "OPTIONS" {
		log.Println("Get request successfull")
		return
	}
}
