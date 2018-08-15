package main

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	//Get the connection string required, function not on github
	sqlInfo := a.GetSQLInfo()
	//Connect to postgres server
	var err error
	a.DB, err = sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	//Ping to ensure connection is valid
	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {

}
