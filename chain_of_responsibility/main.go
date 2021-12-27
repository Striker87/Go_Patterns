package main

import "patterns/chain_of_responsibility/pkg"

/*
Цепочка обязанностей - это поведенческий паттерн проектирования, который позволяет передавать задачи на обработку последовательно по цепочке. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше.
*/

func main() {
	device := &pkg.Device{Name: "Device-1"}
	updateService := &pkg.UpdateDataService{Name: "Update-1"}
	dataService := &pkg.DataDervice{}

	device.SetNext(updateService)
	updateService.SetNext(dataService)

	data := &pkg.Data{}
	device.Execute(data)
}
