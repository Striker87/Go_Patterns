package main

import "patterns/composite/pkg"

/*
Компоновщик - структурный паттерн проектирования, который позволяет сгруппировать множество структур в дерево объектов , а затем работать с ней так, как будто это один.
*/

var (
	cpu = pkg.Cpu{
		Name:        "Corei7 5820k",
		Description: "Intel CPU",
	}
	cpu2 = pkg.Cpu{
		Name:        "Corei9 12900k",
		Description: "Intel CPU",
	}
	gpu1 = pkg.Gpu{
		Name:        "nVidia 1050Ti",
		Description: "nVidia GPU",
	}
	gpu2 = pkg.Gpu{
		Name:        "AMD Radeon 4870",
		Description: "ATI GPU",
	}
	motherboard = pkg.Motherboard{
		Name:        "Asus Sabertooth X99",
		Description: "Asus motherboard",
		Components: []pkg.Component{
			cpu, cpu2, gpu1, gpu2,
		},
	}
	pc = pkg.Pc{
		Name:        "Everest",
		Description: "PC",
		Components: []pkg.Component{
			motherboard,
		},
	}
)

func main() {
	pc.Search("Asus Sabertooth X99")
}
