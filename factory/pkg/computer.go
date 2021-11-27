package pkg

import "fmt"

const (
	ServerType     = "server"
	PersonalPcType = "personal"
	NotebookType   = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PersonalPcType:
		return NewPc()
	case NotebookType:
		return NewNotebook()
	default:
		fmt.Printf("%s не существующий тип\n", typeName)
		return nil
	}
}
