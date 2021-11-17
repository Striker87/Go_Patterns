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

func (b Bank) CheckBalance(cardNumber string) error {
	fmt.Printf("[CheckBalance Банк] Получение остатка по карте %s\n", cardNumber)
	time.Sleep(time.Millisecond * 300)

	for _, card := range b.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[CheckBalance Банк] недостаточно средств")
		}
	}
	fmt.Println("[CheckBalance Банк] Остаток положительный!")

	return nil
}
