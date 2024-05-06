package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	d := demo{
		Name:   "demo",
		Author: "rochitsen",
	}

	output, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)	
}

type demo struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}
