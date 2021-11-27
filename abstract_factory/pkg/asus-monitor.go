package pkg

import "fmt"

type AsusMonitor struct {
	Size int
}

func (m AsusMonitor) PrintDetails() {
	fmt.Printf("[Asus] Monitor Size:[%d]\n", m.Size)
}
