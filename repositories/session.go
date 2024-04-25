package repositories

import (
	"database/sql"
	"go-bank/structs"
)

func InsertSession(db *sql.DB, session structs.Session) (structs.Session, error) {
	query := `INSERT INTO sessions (session_id, user_id, token) VALUES ($1, $2, $3) RETURNING token, expires_at`
	err := db.QueryRow(query, session.ID, session.UserID, session.Token).Scan(&session.Token, &session.ExpiresAt)
	if err != nil {
		return session, err
	}

	return session, nil
}
