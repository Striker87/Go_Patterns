package main

import (
	"fmt"

	"patterns/builder3_director/computer"
)

func main() {
	officeCompBuilder := computer.NewOfficeComputerBuilder()
	officeCompBuilder = officeCompBuilder.RAM(32)
	director := computer.NewDirector(officeCompBuilder)
	compB := director.BuildComputer()
	fmt.Println(compB)
}
