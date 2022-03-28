package pkg

import "fmt"

type Subscriber struct {
	Name string
}

func (s *Subscriber) Update(pubName string) {
	fmt.Printf("Sending to subscribe %s from piblisher %s\n", s.GetName(), pubName)
}

func (s *Subscriber) GetName() string {
	return s.Name
}
