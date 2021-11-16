package main

import "patterns/bridge/pkg/types"

/*
Мост - структурный паттерн проектирования, который разделяет один или несколько объектов на две отдельные абстракцию и реализацию, позволяя изменять их независимо друг от друга.
*/

var (
	hpScanner    = types.HP{Name: "HP ScanJet Pro 2500"}
	epsonScanner = types.Epson{Name: "Epson Perfection V39"}

	linuxPc   = types.LinuxPc{}
	windowsPc = types.WindowsPc{}
	macPc     = types.MacPc{}
)

func main() {
	linuxPc.AddScanner(hpScanner)
	linuxPc.Scan()

	linuxPc.AddScanner(epsonScanner)
	linuxPc.Scan()

	windowsPc.AddScanner(epsonScanner)
	windowsPc.Scan()

	macPc.AddScanner(hpScanner)
	macPc.Scan()
}
