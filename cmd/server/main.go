package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Aminkay95/Payment-Orchestration-Layer/internal/config"
	"github.com/joho/godotenv"
)

func main() {

	//intialize app
	app := &config.Application{
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime),
	}

	//initialize new Mux
	mux := http.NewServeMux()

	//Mux to handle our routes
	mux.HandleFunc("/", app.HelloWeb)

	//Load port from .env
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	PORT := os.Getenv("PORT")

	//Use our custom Logger to give logs
	app.InfoLog.Printf("Starting server at port: %v", PORT)

	//Spin up server

	error := http.ListenAndServe(PORT, mux)

	//Custom Logs error
	app.ErrorLog.Fatal(error)

}
