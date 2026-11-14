package transactions

import (
	"database/sql"
	"time"
)

// Type to hold an individual transaction in Go memory
type Transaction struct {
	Id               int       `json:"id"`
	MerchantId       int       `json:"merchantId"`
	MerchantOrderId  int       `json:"merchantOrderId"`
	AmountCents      int       `json:"amountCents"`
	GatewayId        int       `json:"gatewayId"`
	GatewayAccountId int       `json:"gatewayAccountId"`
	CardTokenId      int       `json:"cardTokenId"`
	State            string    `json:"state"`
	StateUpdatedAt   time.Time `json:"stateUpdatedAt"`
	CreatedAt        time.Time `json:"createdAt"`
	IsDeleted        bool      `json:"isDeleted"`
}

type TransactionModel struct {
	DB *sql.DB
}

// Create Transaction
func (m *TransactionModel) Create(t *Transaction) (int, error) {

	stmt := `INSERT INTO 
				transactions(merchantId, merchantOrderId, amountcents, gatewayId, gatewayAccountId, cardTokenId, state, stateUpdatedAt, createdAt, isDeleted)
				VALUES(?,?,?,?,?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP(),FALSE) RETURNING id`

	row := m.DB.QueryRow(stmt)

	var transactionId int
	err := row.Scan(&transactionId)

	if err != nil {
		return 0, err
	}

	return transactionId, err
}
