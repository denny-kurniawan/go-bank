package repositories

import (
	"database/sql"
	"go-bank/structs"
)

func InsertUser(db *sql.DB, user structs.User) (structs.UserResponse, error) {
	sql := "INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING user_id, username, created_at, updated_at"

	var userResponse structs.UserResponse
	err := db.QueryRow(sql, user.Username, user.Password).Scan(&userResponse.ID, &userResponse.Username, &userResponse.CreatedAt, &userResponse.UpdatedAt)

	return userResponse, err
}

func GetUserByUsername(db *sql.DB, loginInfo structs.LoginInfo) (structs.User, error) {
	sql := "SELECT user_id, username, password_hash, created_at, updated_at FROM users WHERE username=$1"

	var user structs.User
	err := db.QueryRow(sql, loginInfo.Username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}

func UpdatePassword(db *sql.DB, changePasswordInfo structs.ChangePasswordInfo) error {
	sql := "UPDATE users SET password_hash=$1 WHERE username=$2"

	err := db.QueryRow(sql, changePasswordInfo.NewPassword, changePasswordInfo.Username)

	return err.Err()
}

func DeleteUser(db *sql.DB, loginInfo structs.LoginInfo) error {
	sql := "DELETE FROM users WHERE username=$1"

	err := db.QueryRow(sql, loginInfo.Username)

	return err.Err()
}
