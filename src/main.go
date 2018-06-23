package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("home.html"))
	data := HomeModel{Title: "What do you want to watch?"}
	tmpl.Execute(w, data)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("results.html"))
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"term": r.URL.Query().Get("q"),
		}).
		SetHeader("Accept", "application/json").
		SetHeader("X-Mashape-Key", os.Getenv("APIKEY")).
		Get("https://utelly-tv-shows-and-movies-availability-v1.p.mashape.com/lookup?country=us&term={term}")

	var apiResponse APIResponse
	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(apiResponse.Results))
	for _, element := range apiResponse.Results {
		fmt.Println(element.Name)
	}

	tmpl.Execute(w, apiResponse)
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/results", SearchHandler)
	http.ListenAndServe(":8000", nil)
}
