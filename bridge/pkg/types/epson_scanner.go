package types

import "fmt"

type Epson struct {
	Name string
}

func (s Epson) ScanFile() {
	fmt.Println("Epson scan file")
}
