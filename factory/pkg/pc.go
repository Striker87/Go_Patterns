package pkg

import "fmt"

type Pc struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPc() Computer {
	return Pc{
		Type:    PersonalPcType,
		Core:    4,
		Memory:  16,
		Monitor: true,
	}
}

func (p Pc) GetType() string {
	return p.Type
}

func (p Pc) PrintDetails() {
	fmt.Printf("%s Core:[%d] Memory:[%d] Monitor:[%v]\n", p.Type, p.Core, p.Memory, p.Monitor)
}
