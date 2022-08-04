package pkg

import (
	"fmt"
	"time"
)

type Card struct {
	Name    string
	Balance int
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для получения остатка")
	time.Sleep(500 * time.Millisecond)
	return card.Bank.CheckBalance(card.Name)
}
