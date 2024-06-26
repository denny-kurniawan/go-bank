package structs

import "time"

type Session struct {
	ID        string    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
