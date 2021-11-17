package main

import (
	"fmt"
	"patterns/facade/pkg"
)

/*
Фасад - структурный паттерн проектирования, у которого есть простой интерфейс поведения к сложной системе объектов, фукнций или фреймворку.
*/

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
		Name: "Vasya",
		Card: &card1,
	}
	user2 = pkg.User{
		Name: "Petya",
		Card: &card2,
	}
	prod = pkg.Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = pkg.Shop{
		Name:     "Rozetka",
		Products: []pkg.Product{prod},
	}
)

func main() {
	fmt.Printf("[Банк] Выпуск карт\n")
	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("[%s]\n", user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("\n[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
