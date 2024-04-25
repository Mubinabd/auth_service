package models

type Card struct {
	Id      int
	Name    string
	Bank    string
	Num     string
	Pass    int
	Balance int
}

func NewCard() []*Card {
	return []*Card{
		{
			1,
			"Humo SQB",
			"SQB",
			"1234567890123456",
			1111,
			10000000,
		},
		{
			2,
			"UZCARD Ipak yuli",
			"Ipak yuli",
			"8600140290123456",
			7777,
			10000000,
		},
	}
}
