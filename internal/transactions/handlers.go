package transactions

import (
	"fmt"
	"net/http"
	"strconv"
)

type currentTransaction transaction

func createTransaction(w http.ResponseWriter, r *http.Request) {

	//check method is post to ensure idempotency
	if r.Method != http.MethodPost {

		//Letting the user know which method is allowed
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	}

	//Fake data
	currentTransaction{

		merchantId:      1,
		merchantOrderId: 1,
		amountcents:     100,
		gatewayId:       1,
		cardTokenId:     2,
		state:           "Created",
	}

	transactionId, err := currentTransaction.Create()

	//continue creating transaction
	fmt.Println(("transaction Created"))
}

func viewTransactionById(w http.ResponseWriter, r *http.Request) {

	//Get id from parameter
	idStr := r.URL.Query().Get("id")

	//sanitize it ( convert to int, check if its greater than 0)
	id, err := strconv.Atoi(idStr)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//Continue with transactions function
	fmt.Fprintf(w, "Transtion of id %v fetched", id)

}
