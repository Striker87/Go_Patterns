package pkg

import "fmt"

type Motherboard struct {
	Name        string
	Description string
	Components  []Component
}

func (m Motherboard) GetName() string {
	return m.Name
}

func (m Motherboard) Search(name string) {
	if m.Name == name {
		fmt.Printf("Component [%s] found %s", m.Name, m.Description)
		return
	}

	for _, component := range m.Components {
		component.Search(name)
	}
}
