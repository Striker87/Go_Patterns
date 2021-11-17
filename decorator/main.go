package main

import (
	"fmt"
	"patterns/decorator/pkg"
)

/*
Декоратор — это структурный паттерн проектирования, который позволяет динамически добавлять объектам новую функциональность, оборачивая их в полезные «wrapper».
Разберем реализацию на примере сборки компьютеров, которые зависят от базовой комплектации.
*/

var (
	base = pkg.BasePC{}
	home = pkg.HomePC{
		Cpu:     4,
		Gpu:     1,
		Wrapper: base,
	}
	server = pkg.ServerPC{
		Cpu:     16,
		Memory:  256,
		Wrapper: base,
	}
)

func main() {
	fmt.Printf("Base Price: [%f]\n", base.GetPrice())
	fmt.Printf("Home Cpu: [%d] Gpu: [%d] Price: [%f]\n", home.Cpu, home.Gpu, home.GetPrice())
	fmt.Printf("Server Cpu: [%d] Memory: [%d] Price: [%f]\n", server.Cpu, server.Memory, server.GetPrice())
}
