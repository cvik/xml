package xml

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

const yaml = `abc:
  attributes:
  - xmlns-x: myns
  - attr: 'hej'< då
  children:
  - a:
    attributes:
    - hello: pumpkin
    children:
  - b:
    attributes:
    - x-name: me
    children:
  - c:
    attributes:
    children:
  - d:
    attributes:
    children:
    - e:
      attributes:
      children:
    - f:
      attributes:
      children:
      - g:
        attributes:
        children:
      - h:
        attributes:
        children:
`

func TestElement(t *testing.T) {
	el := NewElement()
	if el.Namespaces == nil {
		t.Error("Failed to create new element")
	}

	t.Run("NewElement", func(t *testing.T) { testNewElement(t) })
	t.Run("ParseElement", func(t *testing.T) { testParseElement(t) })
	t.Run("GetAttributeValue", func(t *testing.T) { testGetAttribute(t) })
	t.Run("ToYaml", func(t *testing.T) { testToYaml(t) })
}

// Subtests -------------------------------------------------------------------

func testNewElement(t *testing.T) {
	el := NewElement()
	if el.Namespaces == nil {
		t.Error("Failed to create new element")
	}
}

func testParseElement(t *testing.T) {
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

func testGetAttribute(t *testing.T) {
	data := []byte(`<tnr:elem1 xmlns:p1="http://stuff.it" hell:o="world"/>`)

	el := parseElement(data)

	for _, key := range el.GetAttributeKeys() {
		switch {
		case key.Prefix == "xmlns" && key.Name == "p1":
			continue
		case key.Prefix == "hell" && key.Name == "o":
			continue
		default:
			t.Error("Failed to get attribute keys")
		}
	}

	if val, ok := el.GetAttributeValue("p1"); !ok || val != "http://stuff.it" {
		t.Error("Failed to get attribute values")
	}

	if val, ok := el.GetAttributeValue("o"); !ok || val != "world" {
		t.Error("Failed to get attribute values")
	}

	if _, ok := el.GetAttributeValue("nope"); ok {
		t.Error("Managed to get a none-existing value")
	}
}

func testToYaml(t *testing.T) {
	fd, err := os.Open("ex/weird.xml")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := NewParser(fd).ParseXML()
	if err != nil {
		t.Fatal(err)
	}

	out := doc.Root.ToYaml()
	if out != yaml {
		t.Error("Failed to output element-tree as YAML")
	}

}

func testElementQueries(t *testing.T) {
	fd, err := os.Open("ex/weird.xml")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := NewParser(fd).ParseXML()
	if err != nil {
		t.Fatal(err)
	}

	if !doc.Root.Children[0].HasParent("abc") {
		t.Errorf("Could not get element parent")
	}
}
