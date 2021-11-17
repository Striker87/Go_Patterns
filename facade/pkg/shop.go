package pkg

import (
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (s Shop) Sell(user User, product string) error { // фасад
	fmt.Println("[Shop] Запрос к пользователю для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	fmt.Printf("[Sell Магазин] Проверка - может ли пользователь [%s] купить товар\n", user.Name)
	time.Sleep(time.Millisecond * 500)

	for _, prod := range s.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return fmt.Errorf("[Sell Магазин] Недостаточно средств для покупки товара у пользователя %s\n", user.Name)
		}
		fmt.Printf("[Sell Магазин] Товар [%s] куплен\n", prod.Name)
	}

	return nil
}
