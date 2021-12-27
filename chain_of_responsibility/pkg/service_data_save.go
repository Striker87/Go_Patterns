package pkg

import "fmt"

type DataDervice struct {
	Next Service
}

func (d *DataDervice) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("Data not updated")
		return
	}
	fmt.Println("Data saved")

}

func (d *DataDervice) SetNext(service Service) {
	d.Next = service
}
