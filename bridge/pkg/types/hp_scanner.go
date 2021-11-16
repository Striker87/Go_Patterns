package types

import "fmt"

type HP struct {
	Name string
}

func (s HP) ScanFile() {
	fmt.Println("HP scan file")
}
