package pkg

import (
	"fmt"
	"time"
)

type Card struct {
	Name    string
	Bank    *Bank
	Balance float64
}

func (c Card) CheckBalance() error {
	fmt.Println("[CheckBalance] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)

	return c.Bank.CheckBalance(c.Name)
}
