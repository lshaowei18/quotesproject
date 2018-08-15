package main

import (
	"database/sql"
	"errors"
)

type quote struct {
	ID     int
	quote  string
	author string
	genre  []string
	source string
}

func (q *quote) getQuote(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (q *quote) updateQuote(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (q *quote) deleteQuote(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (q *quote) createQuote(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getQuotes(db *sql.DB, start, count int) ([]quote, error) {
	return nil, errors.New("Not implemented")
}
