package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	resty "gopkg.in/resty.v1"
)

type Result struct {
	Name     string    `json:"name,omitempty"`
	Picture  string    `json:"picture,omitempty"`
	Location *Location `json:"locations,omitempty"`
}

type Location struct {
	LocationName string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
}

var results []Result

func Results(w http.ResponseWriter, req *http.Request) {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"term": req.URL.Query().Get("q"),
		}).
		SetHeader("Accept", "application/json").
		SetHeader("X-Mashape-Key", os.Getenv("APIKEY")).
		Get("https://utelly-tv-shows-and-movies-availability-v1.p.mashape.com/lookup?country=us&term={term}")

	fmt.Println(resp, err)
}

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "project")
	flag.Parse()
	router := mux.NewRouter()

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	router.HandleFunc("/results.html", Results).Methods("GET")
	log.Fatal(srv.ListenAndServe())
}
