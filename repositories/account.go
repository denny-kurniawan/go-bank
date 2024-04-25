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

func GetAccountByAccountNo(db *sql.DB, account structs.Account) (structs.Account, error) {
	sql := "SELECT * FROM accounts WHERE account_number=$1 AND user_id=$2"

	err := db.QueryRow(sql, account.AccountNo, account.UserID).Scan(&account.ID, &account.UserID, &account.AccountNo, &account.Balance, &account.CreatedAt, &account.UpdatedAt)

	return account, err
}

func InsertAccount(db *sql.DB, account structs.Account) (structs.Account, error) {
	sql := "INSERT INTO accounts (user_id, account_number) VALUES ($1, $2) RETURNING *"

	err := db.QueryRow(sql, account.UserID, account.AccountNo).Scan(&account.ID, &account.UserID, &account.AccountNo, &account.Balance, &account.CreatedAt, &account.UpdatedAt)

	return account, err
}

func UpdateAccountBalance(db *sql.DB, account structs.Account) error {
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
