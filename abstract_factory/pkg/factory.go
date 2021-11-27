package pkg

import (
	"fmt"
)

type Factory interface {
	GetComputer() Computer
	GetMonitor() Monitor
}

func GetFactory(brand string) (Factory, error) {
	switch brand {
	case Asus:
		return &AsusFactory{}, nil
	case Hp:
		return &HpFactory{}, nil
	default:
		return nil, fmt.Errorf("Производитель %s - не найден", brand)
	}
}
