package types

import (
	"fmt"
	"patterns/bridge/pkg/base"
)

type MacPc struct {
	scanner base.Scanner
}

func (pc *MacPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *MacPc) Scan() {
	fmt.Println("Scan PC Mac system")
	pc.scanner.ScanFile()
}
