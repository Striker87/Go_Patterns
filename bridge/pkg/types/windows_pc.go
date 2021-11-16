package types

import (
	"fmt"
	"patterns/bridge/pkg/base"
)

type WindowsPc struct {
	scanner base.Scanner
}

func (pc *WindowsPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *WindowsPc) Scan() {
	fmt.Println("Scan PC windows system")
	pc.scanner.ScanFile()
}
