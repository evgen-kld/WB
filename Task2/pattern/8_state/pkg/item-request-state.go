package pkg

import "fmt"

type ItemRequestState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestState) RequestItem() error {
	return fmt.Errorf("[ItemRequestState] Item already requested")
}

func (i *ItemRequestState) AddItem(count int) error {
	return fmt.Errorf("[ItemRequestState] Item Dispense in progress")
}

func (i *ItemRequestState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("[ItemRequestState] Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("[ItemRequestState] Money entred is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestState) DispenseItem() error {
	return fmt.Errorf("[ItemRequestState] Please insert money first")
}
