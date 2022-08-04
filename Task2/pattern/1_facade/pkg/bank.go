package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (bank Bank) CheckBalance(cardNumber string) error {
	fmt.Printf("[Банк] Получение остатка по карте %v \n", cardNumber)
	time.Sleep(500 * time.Millisecond)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	fmt.Println("[Банк] Баланс положительный")
	return nil
}
