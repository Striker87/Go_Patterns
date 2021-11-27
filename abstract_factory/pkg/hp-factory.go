package pkg

type HpFactory struct{}

func (f HpFactory) GetComputer() Computer {
	return HpComputer{
		Memory: 16,
		Cpu:    6,
	}
}

func (f HpFactory) GetMonitor() Monitor {
	return HpMonitor{
		Size: 27,
	}
}
