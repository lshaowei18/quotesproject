package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS testquotes
(
id SERIAL,
quote TEXT NOT NULL,
author VARCHAR(50) NOT NULL,
genre VARCHAR(50) NOT NULL,
source VARCHAR(99) DEFAULT '',
CONSTRAINT testquotes_pkey PRIMARY KEY (id)
)`

var a App

func TestMain(m *testing.M) {
	//Initialize and start app
	a = App{}
	a.Initialize()

	//Check if table exist, if not create it
	ensureTableExists()

	//Run test
	code := m.Run()

	//Clear the table
	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM testquotes")
	a.DB.Exec("ALTER SEQUENCE testquotes_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	//Delete all records from the testquotes table
	clearTable()
	//Make get request and execute it
	req, _ := http.NewRequest("GET", "/quotes", nil)
	response := executeRequest(req)
	checkResponseCode(t, response.Code, http.StatusOK)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Want an empty array, got %v", body)
	}
}

func TestGetNonExistentQuote(t *testing.T) {
	//Delete all records from testquotes table
	clearTable()
	//Create new get request and execute for nonexistent quote
	req, _ := http.NewRequest("GET", "/quote/11", nil)
	response := executeRequest(req)
	checkResponseCode(t, response.Code, http.StatusNotFound)

	var got map[string]string
	json.Unmarshal(response.Body.Bytes(), &got)
	want := "Quote not found"
	if got["error"] != want {
		t.Errorf("Expected the 'error' key of the response to be set to %s, got %s", want, got["error"])
	}
}

func TestCreateQuote(t *testing.T) {
	clearTable()

	//Create payload which is a quote
	payload := []byte(`{"quote":"test quote","author":"test author","genre":["test"]}`)

	//Create new post request and executeRequest
	req, _ := http.NewRequest("POST", "/quote", bytes.NewBuffer(payload))
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
