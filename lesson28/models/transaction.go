package models

type Transaction struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Type        string `json:"type"`
	Description string `json:"description"`
	FromCard    string `json:"from_card"`
	ToCard      string `json:"to_card"`
	PaymentName string `json:"PaymentName"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

type GetAllTransResp struct {
	Trans *[]Transaction `json:"trans"`
	Count int
}

type GetBalanceResp struct {
	
}