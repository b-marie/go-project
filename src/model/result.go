package model

type Result struct {
	Name     string     `json:"name,omitempty"`
	Picture  string     `json:"picture,omitempty"`
	Location []Location `json:"locations,omitempty"`
}

type Location struct {
	LocationName string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
}

var results []Result
