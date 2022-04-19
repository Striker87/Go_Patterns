package pkg

import "fmt"

type Cargo struct {
	Dispatcher
}

func (c *Cargo) PermitArrive() {
	fmt.Println("Грузовик: погрузка")
	c.Arrive()
}

func (c *Cargo) Go() {
	fmt.Println("Грузовик: успешно погружен!")
	c.Dispatcher.NotifyAboutGo()
}

func (c *Cargo) Arrive() {
	if !c.CanArrive(c) {
		fmt.Println("Грузовик: погрузка запрещена! На платформе пассажиры")
		return
	}

	fmt.Println("Грузовик: отправлен!")
}
