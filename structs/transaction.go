package structs

type Transaction struct {
	ID              uint64  `json:"id"`
	AccountNo       string  `json:"account_no"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	CreatedAt       string  `json:"created_at"`
}

type TransactionInfo struct {
	AccountNo   string  `json:"account_no"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type TransferInfo struct {
	FromAccountNo string  `json:"from_account_no"`
	ToAccountNo   string  `json:"to_account_no"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
}
