package main

import (
	"database/sql"
	"fmt"
)

type App struct {
	config Config
	db     *sql.DB
}

type Quote struct {
	id     int
	quote  string
	author string
	genre  []uint8
	source string
}

func main() {
	//Create App
	app := App{}
	app.GetConfig()
	//Handle Func
	// http.HandleFunc("/", Handler)
	// //Create and run server
	// err := http.ListenAndServe(":5050", nil)
	// if err != nil {
	// 	log.Fatal("Listen and serve : ", err)
	// }

	app.ConnectToDB()
	quotes := app.GetByGenre("strategy")
	for _, quote := range quotes {
		fmt.Println(quote.quote)
	}
}
