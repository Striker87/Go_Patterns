package pkg

import "fmt"

type HpMonitor struct {
	Size int
}

func (m HpMonitor) PrintDetails() {
	fmt.Printf("[Asus] Monitor Size:[%d]\n", m.Size)
}
