package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	resty "gopkg.in/resty.v1"
)

type HomeModel struct {
	Title string
}

type Result struct {
	Name     string     `json:"name"`
	Picture  string     `json:"picture"`
	Location []Location `json:"locations"`
}

type Location struct {
	LocationName string `json:"name"`
	URL          string `json:"url"`
}

type APIResponse struct {
	Results []Result `json:"results"`
}

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//Render home page template
	tmpl := template.Must(template.ParseFiles("home.html"))

	//Pass in data for home page
	data := HomeModel{Title: "What do you want to watch?"}

	//Home page error handling
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	//Render results page template
	tmpl := template.Must(template.ParseFiles("results.html"))

	//Use resty to query Utelly API
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"term": r.URL.Query().Get("q"),
		}).
		SetHeader("Accept", "application/json").
		SetHeader("X-Mashape-Key", os.Getenv("APIKEY")).
		Get("https://utelly-tv-shows-and-movies-availability-v1.p.mashape.com/lookup?country=us&term={term}")

	//Build API response
	var apiResponse APIResponse

	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

	e := tmpl.Execute(w, apiResponse)
	if e != nil {
		log.Fatalf("execution failed: %s", e)
	}
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/results", SearchHandler)
	http.ListenAndServe(":8000", nil)
}
