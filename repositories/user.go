package repositories

import (
	"database/sql"
	"go-bank/structs"
)

func GetUsers(db *sql.DB) (users []structs.User, err error) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return
}

func InsertUser(db *sql.DB, user structs.User) (structs.User, error) {
	sql := "INSERT INTO users (username, password_hash, created_at) VALUES ($1, $2, NOW()) RETURNING *"

	err := db.QueryRow(sql, user.Username, user.Password).Scan(&user.ID, &user.Username, &user.CreatedAt)

	return user, err
}

func UpdateUser(db *sql.DB, user structs.User) (structs.User, error) {
	sql := "UPDATE users SET usernam=$1WHERE id=$2 RETURNING *"

	err := db.QueryRow(sql, user.Username, user.ID).Scan(&user.ID, &user.Username, &user.CreatedAt)

	return user, err
}

func DeleteUser(db *sql.DB, user structs.User) error {
	sql := "DELETE FROM users WHERE id=$1"

	err := db.QueryRow(sql, user.ID)

	return err.Err()
}
