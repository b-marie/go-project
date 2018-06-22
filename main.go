package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/resty.v1"
)

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	resp, err := resty.R().Get("http://httpbin.org/get")

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPersonEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe("12345", router))
}
