package main

import "patterns/fabric/pkg"

/*
Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс поведения для создаваемых объектов и у него есть общий интерфейс для этих объектов //объектов который позволяя подтипам реализовывать свою логику работы.
*/

var types = []string{pkg.PersonalPcType, pkg.NotebookType, pkg.ServerType, "NetTop"}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}

		computer.PrintDetails()
	}
}
