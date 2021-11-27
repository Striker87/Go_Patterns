package pkg

import "fmt"

type HpComputer struct {
	Memory int
	Cpu    int
}

func (pc HpComputer) PrintDetails() {
	fmt.Printf("[HP] Pc CPU:[%d] Mem:[%d]\n", pc.Cpu, pc.Memory)
}
