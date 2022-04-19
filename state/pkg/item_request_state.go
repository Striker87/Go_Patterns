package pkg

import "fmt"

type ItemRequestState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *ItemRequestState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. please, insert %d", money)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestState) DispenseItem() error {
	return fmt.Errorf("please insert monet first")
}
