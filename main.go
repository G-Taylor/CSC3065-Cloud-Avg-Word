package main

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error    bool   `json:"error"`
	Sentence string `json:"sentence"`
	Answer   int    `json:"answer"`
}

func queryHandler(w http.ResponseWriter, r *http.Request) {

	inputString := r.URL.Query().Get("text")

	outputJSON := response{
		Error:    false,
		Sentence: "",
		Answer:   0,
	}

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if inputString == "" {
		outputJSON.Error = true
		outputJSON.Sentence = "No Text Entered"
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outputJSON)
}
