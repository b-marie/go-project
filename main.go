package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, req *http.Request) {

}

func Results(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(results)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/results", Results).Methods("GET")
	log.Fatal(http.ListenAndServe("12345", router))
}
