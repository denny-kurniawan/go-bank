package repositories

import (
	"database/sql"
	"go-bank/structs"
)

func InsertTransaction(db *sql.DB, transaction structs.Transaction) (structs.Transaction, error) {
	sql := "INSERT INTO transactions (account_number, transaction_type, amount, description) VALUES ($1, $2, $3, $4) RETURNING *"

	err := db.QueryRow(sql, transaction.AccountNo, transaction.TransactionType, transaction.Amount, transaction.Description).Scan(&transaction.ID, &transaction.AccountNo, &transaction.TransactionType, &transaction.Amount, &transaction.Description, &transaction.CreatedAt)

	return transaction, err
}
