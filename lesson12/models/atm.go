package models

type ATM struct {
	Id           int
	Bank         string
	Location     string
	Money        int
	IsNFCContact bool
}

type ATMBuilder interface {
	CheckPass(p int, c Card) bool
	CheckBalance(c Card) int
	WithDraw(sum int, c Card) bool
	ChangePass(newPass1, newPass2 int, c *Card) bool
}

func NewATM() ATMBuilder {
	return &ATM{1, "SQB", "Qatortol", 10000000, true}
}

func (a *ATM) CheckPass(p int, c Card) bool {
	return c.Pass == p
}

func (a *ATM) CheckBalance(c Card) int {
	return c.Balance
}

func (a *ATM) WithDraw(sum int, c Card) bool {
	if c.Balance >= sum {
		c.Balance -= sum
		a.Money -= sum

		return true
	}

	return false
}

func (a *ATM) ChangePass(newPass1, newPass2 int, c *Card) bool {
	if newPass1 != newPass2 {
		return false
	}

	c.Pass = newPass1

	return true
}
