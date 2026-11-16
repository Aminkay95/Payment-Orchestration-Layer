package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Aminkay95/Payment-Orchestration-Layer/internal/app"
	"github.com/Aminkay95/Payment-Orchestration-Layer/internal/database"
)

func main() {

	PORT := os.Getenv("SERVER_PORT")

	//Connection to database
	db, err := database.NewConnectionDB()

	//intialize app
	app := &app.Application{
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime),
	}

	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	//defer call to db.Close() so that connection is closed before main() exits
	defer db.Close()

	//Use our custom Logger to give logs
	app.InfoLog.Printf("Starting server at port: %v", PORT)

	//Spin up server

	error := http.ListenAndServe(":"+PORT, app.Routes())

	//Custom Logs error
	app.ErrorLog.Fatal(error)

}
