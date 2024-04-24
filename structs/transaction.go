package structs

type Transaction struct {
	ID              uint64  `json:"id"`
	AccountID       uint64  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	CreatedAt       string  `json:"created_at"`
}
