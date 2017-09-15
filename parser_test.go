package xml

import (
	"os"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseXML(t *testing.T) {
	//fd, err := os.Open("ex/CampaignService.wsdl")
	fd, err := os.Open("ex/bing_xsd2.xsd")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := NewParser(fd).ParseXML()
	if err != nil {
		t.Fatal(err)
	}

	// Sample test
	for _, el := range doc.Root.Children {
		if el.Name.Name == "complexType" {
			for _, attr := range el.Attributes {
				if attr.name.Name == "name" && attr.value == "AdApiError" {
					elem := el.Children[0].Children[0]
					if elem.Attributes[1].value == "Code" {
						return
					}
				}
			}
		}
	}

	t.Errorf("Failed to parse XML doc")
}

func TestParseProlog(t *testing.T) {
	xmlString := `
		<?xml version="1.0" encoding="UTF-8" ?>
		<doc attr:key="attr_value"/>`
	doc, err := NewParser(strings.NewReader(xmlString)).ParseXML()
	if err != nil {
		t.Fatal(err)
	}

	if doc.Version != "1.0" || doc.Encoding != "UTF-8" {
		t.Error("Failed to parse prolog")
	}
}

func TestParseName(t *testing.T) {
	name := []byte("pref__:name-_op78_")
	qname := parseQName(name)

	if qname.Prefix != "pref__" || qname.Name != "name-_op78_" {
		t.Error("Failed to parse qualified name")
	}
}

func TestParseAttributes(t *testing.T) {
	data := []byte(`xmlns:q1="val1" unqual='hello "there"!' common="biz">`)
	expQnames := []QName{{"xmlns", "q1"}, {"", "unqual"}, {"", "common"}}
	expValues := []string{"val1", `hello "there"!`, "biz"}

	attrs := parseAttributes(data)

	for i, attr := range attrs {
		if attr.name != expQnames[i] || attr.value != expValues[i] {
			t.Error("failed to read attribute with qualified name")
		}
	}
}

func TestParseElement(t *testing.T) {
	data := [][]byte{
		[]byte(`<tnr:elem1 xmlns:p1="http://stuff.it" hell:o="world"/>`),
		[]byte(`</tnr:elem1>`),
		[]byte(`<a>`),
		[]byte(`<b/>`),
	}

	for _, d := range data {
		el := parseElement(d)
		if el == nil {
			t.Error("Failed to parse element")
		}

		if el.Name.Name == "" {
			spew.Dump(el)
		}
	}
}
