package structs

type Account struct {
	ID        uint64  `json:"account_id"`
	UserID    uint64  `json:"user_id"`
	AccountNo string  `json:"account_no"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type AccountBalance struct {
	AccountNo string  `json:"account_no"`
	Balance   float64 `json:"balance"`
}

type AccountTransaction struct {
	ID              uint64  `json:"id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	CreatedAt       string  `json:"created_at"`
}

type AccountDetails struct {
	ID           uint64               `json:"account_id"`
	UserID       uint64               `json:"user_id"`
	AccountNo    string               `json:"account_no"`
	Balance      float64              `json:"balance"`
	CreatedAt    string               `json:"created_at"`
	UpdatedAt    string               `json:"updated_at"`
	Transactions []AccountTransaction `json:"transactions"`
}
