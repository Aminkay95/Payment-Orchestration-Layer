package transactions

import (
	"database/sql"
	"time"
)

// Type to hold an individual transaction
type transaction struct {
	id               int
	merchantId       int
	merchantOrderId  int
	ammountCents     int
	gatewayId        int
	gatewayAccountId int
	cardTokenId      int
	state            string
	stateUpdatedAt   time.Time
	createdAt        time.Time
	isDeleted        bool
}

type TransactionModel struct {
	DB *sql.DB
}

// Create Transaction
func (m *TransactionModel) Create(merchantId, merchantOrderId, ammountCents, gatewayId, gatewayAccountId, cardTokenId int, state string) (int, error) {

	stmt := `INSERT INTO 
				transactions(merchantId, merchantOrderId, amountcents, gatewayId, gatewayAccountId, cardTokenId, state, stateUpdatedAt, createdAt, isDeleted)
				VALUES(?,?,?,?,?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP(),FALSE) RETURNING id`

	row := m.DB.QueryRow(stmt)

	var transactionId int
	err := row.Scan(&transactionId)

	if err != nil {
		return 0, err
	}

	return transactionId, nil
}

// Get Transaction by id
func (m *TransactionModel) GetTransactionById(id int)
