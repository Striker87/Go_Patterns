package service

import "fmt"

type AnalystDataService interface {
	SendXmlData()
}

type XmlDocument struct {
}

func (x XmlDocument) SendXmlData() {
	fmt.Println("sending xml data")
}
