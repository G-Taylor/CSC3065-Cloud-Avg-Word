package main

import (
	"40234272/editor-avg-word/function"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAvgWordCount(t *testing.T) {
	answer := function.AverageWordLength("one two one")

	if answer != 3 {
		t.Error("Expected 3, got ", answer)
	}
}

func TestAvgWordCountNoInput(t *testing.T) {
	answer := function.AverageWordLength("")

	if answer != 0 {
		t.Error("Expected 0, got ", answer)
	}
}

func TestAvgWordCountWithNumbers(t *testing.T) {
	answer := function.AverageWordLength("one 5345 two 34 on43e")

	if answer != 3 {
		t.Error("Expected 3, got ", answer)
	}
}

func TestAvgWordCountWithSpecialChars(t *testing.T) {
	answer := function.AverageWordLength("one '/' two []] on-e")

	if answer != 3 {
		t.Error("Expected 3, got ", answer)
	}
}

func TestAvgWordCount_API(t *testing.T) {
	handler := QueryHandler
	req := httptest.NewRequest("GET", "http://frontend.40234272.qpc.hal.davecutting.uk/?text=one%20two%20one", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	json_response := strings.TrimRight(string(body), "\n")

	expectedResult := `{"error":false,"sentence":"one two one","answer":3}`

	if json_response != expectedResult {
		t.Error("Expected ", expectedResult, " got ", resp)
	}
}

func TestAvgWordCountNoInput_API(t *testing.T) {
	handler := QueryHandler
	req := httptest.NewRequest("GET", "http://frontend.40234272.qpc.hal.davecutting.uk/?text=", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	json_response := strings.TrimRight(string(body), "\n")

	expectedResult := `{"error":true,"sentence":"No Text Entered","answer":0}`

	if json_response != expectedResult {
		t.Error("Expected ", expectedResult, " got ", resp)
	}
}

func TestAvgWordCountWithNumbers_API(t *testing.T) {
	handler := QueryHandler
	req := httptest.NewRequest("GET", "http://frontend.40234272.qpc.hal.davecutting.uk/?text=one%20t3wo%20one3%20435%2045", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	json_response := strings.TrimRight(string(body), "\n")

	expectedResult := `{"error":false,"sentence":"one t3wo one3 435 45","answer":3}`

	if json_response != expectedResult {
		t.Error("Expected ", expectedResult, " got ", resp)
	}
}

func TestAvgWordCountWithSpecialChars_API(t *testing.T) {
	handler := QueryHandler
	req := httptest.NewRequest("GET", "http://frontend.40234272.qpc.hal.davecutting.uk/?text=one%20t-=wo%20one.%27%20]%27", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	json_response := strings.TrimRight(string(body), "\n")

	expectedResult := `{"error":false,"sentence":"one t-=wo one.' ]'","answer":3}`

	if json_response != expectedResult {
		t.Error("Expected ", expectedResult, " got ", resp)
	}
}

func TestAvgWordCountWithBadParameters_API(t *testing.T) {
	handler := QueryHandler
	req := httptest.NewRequest("GET", "http://frontend.40234272.qpc.hal.davecutting.uk/test", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	json_response := strings.TrimRight(string(body), "\n")

	expectedResult := `404 not found.`

	if json_response != expectedResult {
		t.Error("Expected ", expectedResult, " got ", resp)
	}
}
