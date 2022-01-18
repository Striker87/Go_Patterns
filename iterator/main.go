package main

import (
	"fmt"

	"patterns/iterator/pkg"
)

/*
Итератор - поведенческий паттерн проектирования, который позволяет последовательно обходить переменные не раскрывая их внутреннего представления. Применяется когда есть сложные структуры данных и мы хотим скрыть от клиента детали реализации
*/

var routes = &pkg.Routes{
	Routes: []pkg.Route{
		{Name: "Route 1", TravelTime: 110},
		{Name: "Route 2", TravelTime: 90},
		{Name: "Route 3", TravelTime: 60},
		{Name: "Route 4", TravelTime: 40},
	},
}

func main() {
	for routes.HasNext() {
		route := routes.GetNext()
		fmt.Printf("R:[%s] Time:[%d]\n", route.Name, route.TravelTime)
	}
}
