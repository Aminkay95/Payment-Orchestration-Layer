package config

import (
	"log"
	"net/http"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (Application) HelloWeb(w http.ResponseWriter, r *http.Request) {

	//Return Not Found

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Backend is running"))
}
