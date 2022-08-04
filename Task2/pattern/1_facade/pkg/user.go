package pkg

type User struct {
	Name string
	Card *Card
}

func (user User) GetBalance() int {
	return user.Card.Balance
}
