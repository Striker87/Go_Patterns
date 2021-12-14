package pkg

import "fmt"

type Singleton struct {
	Type string
}

func (s Singleton) Print() {
	fmt.Printf("Type %s", s.Type)
}

func NewSingleton(item *Singleton, typeName string) *Singleton {
	if item == nil {
		return &Singleton{Type: typeName}
	}

	fmt.Printf("Type %s - уже создан\n", item.Type)

	return item
}
