package main

import (
	"fmt"

	"patterns/abstract_factory/pkg"
)

/*
Абстрактная фабрика — это порождающий паттерн проектирования, который позволяет создавать группы связанных объектов, не привязываясь к конкретным типам создаваемых объектов.
*/

var (
	brands = []string{pkg.Asus, pkg.Hp, "Dell"}
)

func main() {
	for _, brand := range brands {
		factory, err := pkg.GetFactory(brand)
		if err != nil {
			fmt.Println(err)
			continue
		}
		monitor := factory.GetMonitor()
		monitor.PrintDetails()

		computer := factory.GetComputer()
		computer.PrintDetails()
	}
}
