package pkg

import "fmt"

type Pc struct {
	Name        string
	Description string
	Components  []Component
}

func (p Pc) GetName() string {
	return p.Name
}

func (p Pc) Search(name string) {
	if p.Name == name {
		fmt.Printf("Component [%s] found %s", p.Name, p.Description)
		return
	}

	for _, component := range p.Components {
		component.Search(name)
	}
}
