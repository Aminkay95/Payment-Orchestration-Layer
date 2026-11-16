package app

import (
	"net/http"

	"github.com/Aminkay95/Payment-Orchestration-Layer/internal/transactions"
)

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	//Api health check
	mux.HandleFunc("GET /", app.HelloWeb)

	transactions.RegisterRoutes(mux)

	return mux

}
