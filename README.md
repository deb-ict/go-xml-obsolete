# go-xml
This package is a copy of the encoding/xml package with namespace prefix support.

# Installation
`go get -u github.com/deb-ict/go-xml`

Replace import `"encoding/xml"` with `"github.com/deb-ict/go-xml"`

# Example
This example show encoding & decoding of a SOAP envelope with namespace prefix support
```
package main

import (
	"fmt"
	"os"

	"github.com/deb-ict/go-xml"
)

type Envelope struct {
	XMLName xml.Name   `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Attrs   []xml.Attr `xml:",any,attr"`
	Header  *Header
	Body    *Body
}

type Header struct {
	XMLName xml.Name   `xml:"http://www.w3.org/2003/05/soap-envelope Header"`
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
}

type Body struct {
	XMLName xml.Name   `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
}

type AlertControl struct {
	XMLName  xml.Name   `xml:"http://example.org/alertcontrol alertcontrol"`
	Attrs    []xml.Attr `xml:",any,attr"`
	Priority int        `xml:"http://example.org/alertcontrol priority"`
	Expires  time.Time  `xml:"http://example.org/alertcontrol expires"`
}

type Alert struct {
	XMLName xml.Name   `xml:"http://example.org/alert alert"`
	Attrs   []xml.Attr `xml:",any,attr"`
	Message string     `xml:"http://example.org/alert msg"`
}

func testXmlIn() {
	xmlFile, _ := os.Open("xml/soap-001.xml")
	xmlEnvelope := &Envelope{}

	decoder := xml.NewDecoder(xmlFile)
	decoder.Decode(&xmlEnvelope)

	fmt.Println(xmlEnvelope.XMLName.Space)
}

func testXmlOut() {
	nametable := make(map[string]string)
	nametable["http://example.org/alertcontrol"] = "ac"
	nametable["http://example.org/alert"] = "a"

	headerContent, _ := xml.MarshalIndentWithNametable(&AlertControl{
		Priority: 1,
		Expires:  time.Now().Add(48 * time.Hour),
	}, nametable, "", "  ")
	bodyContent, _ := xml.MarshalWithNametable(&Alert{
		Message: "Hello alert!",
	}, nametable)

	xmlEnvelope := &Envelope{
		Attrs: []xml.Attr{
			{
				Name:  xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "attr1"},
				Value: "value1",
			},
			{
				Name:  xml.Name{Space: "http://www.w3.org/2003/06/test", Local: "attr1"},
				Value: "value1",
			},
		},
		Header: &Header{
			Content: headerContent,
		},
		Body: &Body{
			Content: bodyContent,
		},
	}

	encoder := xml.NewEncoder(os.Stdout)
	encoder.SetNametable(nametable)
	encoder.SetNamespace("soap", XMLNS_SOAP_12_ENV)
	encoder.Indent("", "  ")
	encoder.Encode(xmlEnvelope)
}

func main() {
	testXmlIn()
	testXmlOut()
}
```

## Output
```
<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" soap:attr1="value1" xmlns:test="http://www.w3.org/2003/06/test" test:attr1="value1">
  <soap:Header><ac:alertcontrol xmlns:ac="http://example.org/alertcontrol">
  <ac:priority>1</ac:priority>
  <ac:expires>2024-12-22T16:00:17.9183821+01:00</ac:expires>
</ac:alertcontrol></soap:Header>
  <soap:Body><a:alert xmlns:a="http://example.org/alert"><a:msg>Hello alert!</a:msg></a:alert></soap:Body>
</soap:Envelope>
```
