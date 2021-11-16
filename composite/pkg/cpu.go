package pkg

import "fmt"

type Cpu struct {
	Name        string
	Description string
}

func (c Cpu) GetName() string {
	return c.Name
}

func (c Cpu) Search(name string) {
	if c.Name == name {
		fmt.Printf("Component [%s] found %s", c.Name, c.Description)
	}
}
