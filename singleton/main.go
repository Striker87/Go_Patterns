package main

import (
	"fmt"
	"patterns/singleton/pkg"
)

/*
Одиночка — это порождающий паттерн проектирования, который гарантирует, что у объекта есть только один экземпляр, и предоставляет к нему глобальный доступ.
*/

var singleton *pkg.Singleton

func init() {
	fmt.Println("init")

	singleton = &pkg.Singleton{Type: "Одиночка"}
}

func main() {
	for i := 0; i < 3; i++ {
		singleton = pkg.NewSingleton(singleton, "Create singleton")

		fmt.Println(singleton)
	}
}
