package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю для получения остатка по карте")
	time.Sleep(500 * time.Millisecond)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли %s купить товар \n", user.Name)
	time.Sleep(500 * time.Millisecond)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара!")
		}
		fmt.Printf("[Магазин] Товар [%s] - куплен \n", prod.Name)
	}
	return nil
}
