package main

import "fmt"

type xmlSender interface {
	sendXml() string
}

type xmlDocument struct {
	body string
}

func (doc xmlDocument) sendXml() string {
	return doc.body
}

type jsonDocument struct {
	data map[string]any
}

func (doc jsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

type jsonDocumentAdapter struct {
	jsonDocument *jsonDocument
}

func (adapter jsonDocumentAdapter) sendXml() string {
	return adapter.jsonDocument.ConvertToXml()
}

func sendToXmlSystem(s xmlSender) {
	fmt.Println("got: ", s.sendXml())
}

func main() {
	xmlDoc := xmlDocument{body: "<></>"}
	jsonDoc := jsonDocument{data: map[string]any{"1": "ok", "2": "not ok"}}

	sendToXmlSystem(xmlDoc)
	sendToXmlSystem(jsonDocumentAdapter{jsonDocument: &jsonDoc})
}
