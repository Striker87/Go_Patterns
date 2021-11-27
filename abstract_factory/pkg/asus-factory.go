package pkg

type AsusFactory struct{}

func (f AsusFactory) GetComputer() Computer {
	return AsusComputer{
		Memory: 8,
		Cpu:    4,
	}
}

func (f AsusFactory) GetMonitor() Monitor {
	return AsusMonitor{
		Size: 32,
	}
}
