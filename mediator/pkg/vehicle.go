package pkg

type Vehicle interface {
	Arrive()
	Go()
	PermitArrive()
}

type Dispatcher interface {
	CanArrive(vehicle Vehicle) bool
	NotifyAboutGo()
}
