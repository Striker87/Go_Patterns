package pkg

import "fmt"

type AsusComputer struct {
	Memory int
	Cpu    int
}

func (pc AsusComputer) PrintDetails() {
	fmt.Printf("[Asus] Pc CPU:[%d] Mem:[%d]\n", pc.Cpu, pc.Memory)
}
