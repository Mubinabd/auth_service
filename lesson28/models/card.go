package models

type Card struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CardNum   string `json:"card_num"`
	ExpiredAt string `json:"expired_at"`
	Amount    int    `json:"amount"`
	Password  int16  `json:"password"`
	CardType  string `json:"card_type"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

type GetAllCardReq struct {
	UserId   string
	CardType string
}

type GetAllCardsResp struct {
	Cards *[]Card
	Count int
}

type UpdatePassReq struct {
	ID   string
	Pass int16
}
