package structs

type Account struct {
	ID        uint64  `json:"id"`
	UserID    uint64  `json:"user_id"`
	AccountNo string  `json:"account_no"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}
