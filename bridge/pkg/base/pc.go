package base

type PC interface {
	Scan()
	AddScanner(scanner Scanner)
}
