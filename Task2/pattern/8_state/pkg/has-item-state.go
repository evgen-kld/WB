package pkg

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("[HasItemState] No Item present")
	}
	fmt.Println("[HasItemState] Item Request")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Println("[HasItemState] Items added")
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("[HasItemState] Please select item first")
}

func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("[HasItemState] Please select item first")
}
