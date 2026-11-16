package app

import (
	"encoding/json"
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

	w.Header().Set("Content-Type", "application/json")

	msg := map[string]any{
		"msg": "Backend is running",
	}

	json.NewEncoder(w).Encode(msg)
}
