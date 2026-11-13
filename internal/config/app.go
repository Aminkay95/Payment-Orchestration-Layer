package config

import (
	"log"
	"net/http"

	"github.com/Aminkay95/Payment-Orchestration-Layer/internal/transactions"
)

type Application struct {
	InfoLog     *log.Logger
	ErrorLog    *log.Logger
	transaction *transactions.TransactionModel
}

func (Application) HelloWeb(w http.ResponseWriter, r *http.Request) {

	//Return Not Found

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Backend is running"))
}
