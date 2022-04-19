package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (m *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (m *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (m *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (m *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing item")
	m.vendingMachine.itemCount = m.vendingMachine.itemCount - 1

	if m.vendingMachine.itemCount == 0 {
		m.vendingMachine.setState(m.vendingMachine.noItem)
	} else {
		m.vendingMachine.setState(m.vendingMachine.hasItem)
	}
	return nil
}
