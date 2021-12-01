package main

import (
	"40234272/editor-avg-word/function"
	"encoding/json"
	"net/http"
)

type Response struct {
	Error    bool    `json:"error"`
	Sentence string  `json:"sentence"`
	Answer   float32 `json:"answer"`
}

func main() {
	http.HandleFunc("/", QueryHandler)
	http.ListenAndServe(":5000", nil)
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {

	inputString := r.URL.Query().Get("text")

	outputJSON := Response{
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
	} else {
		outputJSON.Sentence = inputString
		answer := function.AverageWordLength(inputString)
		outputJSON.Answer = answer
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outputJSON)
}
