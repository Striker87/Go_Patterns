package main

import "patterns/builder/pkg"

/*
Строитель — порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово. Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.
*/

func main() {
	asusCollector := pkg.GetCollector("asus")
	hpCollector := pkg.GetCollector("hp")

	builder := pkg.NewBuilder(asusCollector)
	asusComputer := builder.CreateComputer()
	asusComputer.Print()

	builder.SetCollector(hpCollector)
	hpComputer := builder.CreateComputer()
	hpComputer.Print()
}
