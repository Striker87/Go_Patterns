package pkg

import "fmt"

type Passenger struct {
	Dispatcher
}

func (p *Passenger) PermitArrive() {
	fmt.Println("Пассажиры: занимайте места!")
	p.Arrive()
}

func (p *Passenger) Go() {
	fmt.Println("Пассажиры: отправление!")
	p.Dispatcher.NotifyAboutGo()
}

func (p *Passenger) Arrive() {
	if !p.CanArrive(p) {
		fmt.Println("Пассажиры: отправление задерживается! Платформа занята")
		return
	}

	fmt.Println("Пассажиры: занимайте места!")
}
