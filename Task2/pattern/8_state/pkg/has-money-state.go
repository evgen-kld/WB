package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("[HasMoneyState] Item Dispense in progress")
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("[HasMoneyState] Item Dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("[HasMoneyState] Item out of stock")
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("[HasMoneyState] Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
