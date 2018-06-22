package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

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

// func Home(w http.ResponseWriter, req *http.Request) {
// 	w.Write([]byte("Hello There"))
// }

func Results(w http.ResponseWriter, req *http.Request) {
	resp, err := resty.R().
		// SetQueryParams(map[string]string{
		// 	"term": "search_term",
		// }).
		SetHeader("Accept", "application/json").
		SetAuthToken("P7F18tP3Ulmsh3ZFIv1FyVYygMoBp1b1ZoqjsnIWOkMPL0osxc").
		Get("https://utelly-tv-shows-and-movies-availability-v1.p.mashape.com/lookup?country=us&term=bojack")

	fmt.Println(resp, err)
}

func main() {
	router := mux.NewRouter()
	templates := populateTemplates()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	router.HandleFunc("/results", Results).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
