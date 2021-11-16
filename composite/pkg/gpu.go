package pkg

import "fmt"

type Gpu struct {
	Name        string
	Description string
}

func (g Gpu) GetName() string {
	return g.Name
}

func (g Gpu) Search(name string) {
	if g.Name == name {
		fmt.Printf("Component [%s] found %s", g.Name, g.Description)
	}
}
