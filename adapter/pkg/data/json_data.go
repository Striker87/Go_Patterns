package data

import "fmt"

type JsonDocument struct {
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (j JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

func (a JsonDocumentAdapter) SendXmlData() {
	a.jsonDocument.ConvertToXml()
	fmt.Println("Sending XML data")
}
