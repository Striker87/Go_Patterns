package types

import (
	"fmt"
	"patterns/bridge/pkg/base"
)

type LinuxPc struct {
	scanner base.Scanner
}

func (pc *LinuxPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *LinuxPc) Scan() {
	fmt.Println("Scan PC linux system")
	pc.scanner.ScanFile()
}
