package pkg

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (u *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data from device [%s] already updated\n", u.Name)
		u.Next.Execute(data)
		return
	}

	fmt.Printf("Update Data from device [%s]\n", u.Name)
	data.UpdateSource = true
	u.Next.Execute(data)
}

func (u *UpdateDataService) SetNext(service Service) {
	u.Next = service
}
