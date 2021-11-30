package pkg

import "fmt"

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (c Computer) Print() {
	fmt.Printf("%s Core:[%d] Mem:[%d] GraphicCard:[%d] Monitor:[%d]\n", c.Brand, c.Core, c.GraphicCard, c.Memory, c.Monitor)
}
