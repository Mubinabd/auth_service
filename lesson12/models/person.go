package models

type Person struct {
	Name  string
	Cash  int
	Cards []*Card
}

func NewPerson(c []*Card) Person {
	return Person{"Husan", 100000, c}
}

func (p Person) SelectCard(id int) Card {
	for _, v := range p.Cards {
		if id == v.Id {
			return *v
		}
	}

	return Card{}
}