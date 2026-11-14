package transactions

import (
	"encoding/json"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	//check method is post to ensure idempotency
	if r.Method != http.MethodPost {

		//Letting the user know which method is allowed
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	}

	//Fake data
	currentTransaction := &Transaction{
		MerchantId:      1,
		MerchantOrderId: 1,
		AmountCents:     100,
		GatewayId:       1,
		CardTokenId:     2,
		State:           "Created",
	}

	//send response of the our fake data

	//Setting header
	w.Header().Set("Content-Type", "application/json")

	data := map[string]any{
		"data": currentTransaction,
	}

	json.NewEncoder(w).Encode(data)

}
