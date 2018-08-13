package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func (app *App) ConnectToDB() {
	var err error
	//Open DB
	app.db, err = sql.Open("postgres", app.config.sqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	//Ping to make sure there is connection
	err = app.db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func (app App) GetByGenre(name string) []Quote {
	statement := "SELECT * FROM quotes_v1 WHERE $1 = ANY (genre);"
	var quotes []Quote
	rows, err := app.db.Query(statement, name)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var row Quote
		err := rows.Scan(&row.id, &row.quote, &row.author, &row.genre, &row.source)
		if err != nil {
			log.Fatalln(err)
		}
		quotes = append(quotes, row)
	}
	return quotes
}
