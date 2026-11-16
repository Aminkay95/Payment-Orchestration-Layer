package transactions

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /transactions", handleCreateTransaction)

}
