package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func getHomeRequest(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestHomeHandler(t *testing.T) {
	r := getHomeRequest(t, "/")

	rw := httptest.NewRecorder()

	HomeHandler(rw, r)
}

func BenchmarkAccessHomeRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := getHomeRequest(b, "/")
		rw := httptest.NewRecorder()
		HomeHandler(rw, r)
	}
}

func getResultsRequest(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestSearchHandler(t *testing.T) {
	r := getResultsRequest(t, "/results")

	rw := httptest.NewRecorder()

	SearchHandler(rw, r)
}

func BenchmarkAccessResultsRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := getResultsRequest(b, "/results")
		rw := httptest.NewRecorder()
		SearchHandler(rw, r)
	}
}
