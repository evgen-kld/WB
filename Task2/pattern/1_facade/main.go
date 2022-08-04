package main

import (
	"Task2/pattern/1_facade/pkg"
	"fmt"
)

var (
	bank = pkg.Bank{
		Name:  "Банк",
		Cards: []pkg.Card{},
	}
	card1 = pkg.Card{
		Name:    "Card-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = pkg.Card{
		Name:    "Card-2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 = pkg.User{
		Name: "Покупатель 1",
		Card: &card1,
	}
	user2 = pkg.User{
		Name: "Покупатель 2",
		Card: &card2,
	}
	prod = pkg.Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = pkg.Shop{
		Name: "Магазин",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]", user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("[%s]", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
