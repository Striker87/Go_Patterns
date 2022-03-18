package main

import (
	"fmt"

	"github.com/kr/pretty"
)

// Используется для построения сложных объектов покомпонентно

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI // motherboard

	Build() Computer
}

func (b *computerBuilder) CPU(val string) ComputerBuilderI {
	b.cpu = val
	return b
}
func (b *computerBuilder) RAM(val int) ComputerBuilderI {
	b.ram = val
	return b
}
func (b *computerBuilder) MB(val string) ComputerBuilderI {
	b.mb = val
	return b
}

func (b computerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

func NewComputerBuilder() ComputerBuilderI {
	return &computerBuilder{}
}

func main() {
	computerBuilder := NewComputerBuilder()
	computer := computerBuilder.CPU("Core i3").RAM(8).MB("Asus").Build()

	fmt.Printf("%# v\n\n", pretty.Formatter(computer))

	officeComputerBuilder := NewOfficeComputerBuilder()
	officeComputerBuilder.RAM(4)
	officeComputer := officeComputerBuilder.Build()

	fmt.Printf("%# v\n\n", pretty.Formatter(officeComputer))
}

type officeComputerBuilder struct {
	computerBuilder
}

func NewOfficeComputerBuilder() ComputerBuilderI {
	return &officeComputerBuilder{}
}

func (b *officeComputerBuilder) Build() Computer {
	if b.cpu == "" {
		b.CPU("celeron")
	}
	if b.ram == 0 {
		b.RAM(2)
	}
	if b.mb == "" {
		b.MB("ASRock")
	}

	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}
