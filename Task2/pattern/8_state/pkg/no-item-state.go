package pkg

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("[NoItemState] Item out of stock")
}

func (i *NoItemState) AddItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	fmt.Println("[NoItemState] Items added")
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("[NoItemState] Item out of stock")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("[NoItemState] Item out of stock")
}
