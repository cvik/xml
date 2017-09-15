package xml

import (
	"fmt"
	"strings"
)

type Namespace struct {
	Prefix string
	URI    string
}

// Element --------------------------------------------------------------------

type Element struct {
	Name       QName
	Attributes []attribute
	Namespaces map[string]string
	Children   []*Element
	Parent     *Element
	Closed     bool
	Text       string
}

func NewElement() *Element {
	return &Element{
		Namespaces: make(map[string]string),
	}
}

func (el *Element) GetAttributeValue(name string) (string, bool) {
	for _, attr := range el.Attributes {
		if attr.name.Name == name {
			return attr.value, true
		}
	}
	return "", false
}

func (el *Element) GetAttributeKeys() []QName {
	tmp := []QName{}
	for _, attr := range el.Attributes {
		tmp = append(tmp, attr.name)
	}
	return tmp
}

func (el *Element) HasParent(name string) bool {
	if el.Parent.Name.Name == name {
		return true
	}
	return false
}

func (el *Element) GetChildren(name string) []*Element {
	tmp := []*Element{}
	for _, c := range el.Children {
		if c.Name.Name == name {
			tmp = append(tmp, c)
		}
	}
	return tmp
}

func (el *Element) GetNamespace(name QName) string {
	ns, ok := el.Namespaces[name.Prefix]
	if !ok {
		ns, ok := el.Namespaces["$default"]
		if !ok {
			return ""
		}
		return ns
	}
	return ns
}

func (el *Element) ToYaml() string {
	return el.toYaml(0)
}

// Internal -------------------------------------------------------------------

func (el *Element) toYaml(indent int) string {
	indentStr, pref := strings.Repeat(" ", indent), ""
	if el.Name.Prefix != "" {
		pref = el.Name.Prefix + "-"
	}
	tmp := fmt.Sprintf("%s%s:\n", pref, el.Name.Name)
	tmp += indentStr + "  attributes:\n"
	for _, attr := range el.Attributes {
		pref = ""
		if attr.name.Prefix != "" {
			pref = attr.name.Prefix + "-"
		}
		tmp += fmt.Sprintf("%s  - %s%s: %s\n",
			indentStr, pref, attr.name.Name, attr.value)
	}
	tmp += indentStr + "  children:\n"
	for _, child := range el.Children {
		tmp += indentStr + "  - " + child.toYaml(indent+2)
	}
	return tmp
}

func (el *Element) addNamespaces() {
	tmp := map[string]string{}
	if el.Parent != nil {
		for key, val := range el.Parent.Namespaces {
			tmp[key] = val
		}
	}

	for _, attr := range el.Attributes {
		if strings.ToLower(attr.name.Prefix) == "xmlns" {
			tmp[attr.name.Name] = attr.value
		}
		if attr.name.Prefix == "" && strings.ToLower(attr.name.Name) == "xmlns" {
			tmp["$default"] = attr.value
		}
	}

	el.Namespaces = tmp
}
