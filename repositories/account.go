package repositories

import (
	"database/sql"
	"go-bank/structs"
)

func GetAccountsByUserID(db *sql.DB, userID string) (accounts []structs.Account, err error) {
	sql := "SELECT * FROM accounts WHERE user_id=$1"

	rows, err := db.Query(sql, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account structs.Account
		err := rows.Scan(&account.ID, &account.UserID, &account.AccountNo, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func GetAccountDetails(db *sql.DB, account structs.AccountDetails) (structs.AccountDetails, error) {
	sql1 := "SELECT * FROM accounts WHERE account_number=$1"

	err := db.QueryRow(sql1, account.AccountNo).Scan(&account.ID, &account.UserID, &account.AccountNo, &account.Balance, &account.CreatedAt, &account.UpdatedAt)

	if err != nil {
		return account, err
	}

	sql2 := `SELECT transaction_id, transaction_type, amount, description, created_at FROM transactions WHERE account_number=$1 ORDER BY created_at DESC`

	rows, err := db.Query(sql2, account.AccountNo)
	if err != nil {
		return account, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction structs.AccountTransaction
		err := rows.Scan(&transaction.ID, &transaction.TransactionType, &transaction.Amount, &transaction.Description, &transaction.CreatedAt)
		if err != nil {
			return account, err
		}

		account.Transactions = append(account.Transactions, transaction)
	}

	return account, nil
}

func InsertAccount(db *sql.DB, account structs.Account) (structs.Account, error) {
	sql := "INSERT INTO accounts (user_id, account_number) VALUES ($1, $2) RETURNING *"

	err := db.QueryRow(sql, account.UserID, account.AccountNo).Scan(&account.ID, &account.UserID, &account.AccountNo, &account.Balance, &account.CreatedAt, &account.UpdatedAt)

	return account, err
}

func GetAccountBalance(db *sql.DB, account structs.AccountBalance) (structs.AccountBalance, error) {
	sql := "SELECT account_number, balance FROM accounts WHERE account_number=$1"

	err := db.QueryRow(sql, account.AccountNo).Scan(&account.AccountNo, &account.Balance)

	return account, err
}

func UpdateAccountBalance(db *sql.DB, account structs.AccountBalance) error {
	sql := "UPDATE accounts SET balance=$1, updated_at=NOW() WHERE account_number=$2"

	_, err := db.Exec(sql, account.Balance, account.AccountNo)

	return err
}

func DeleteAccountByAccountNo(db *sql.DB, account structs.Account) error {
	sql := "DELETE FROM accounts WHERE account_number=$1 AND user_id=$2"

	_, err := db.Exec(sql, account.AccountNo, account.UserID)

	return err
}

func DeleteAccountsByUserID(db *sql.DB, userID string) error {
	sql := "DELETE FROM accounts WHERE user_id=$1"

	_, err := db.Exec(sql, userID)

	return err
}
