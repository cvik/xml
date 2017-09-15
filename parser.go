package xml

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Document struct {
	Encoding string
	Version  string
	Root     *Element
}

type Parser struct {
	r          io.Reader
	buffersize int64
}

func NewParser(r io.Reader) *Parser {

	return &Parser{
		r: r,
	}
}

func (p *Parser) ParseXML() (*Document, error) {
	doc := &Document{}
	scanner := bufio.NewScanner(p.r)
	scanner.Split(split)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	current := scanner.Bytes()

	if isProlog(current) {
		el := parseElement(current)
		doc.Version, _ = el.GetAttributeValue("version")
		doc.Encoding, _ = el.GetAttributeValue("encoding")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		current = scanner.Bytes()
	}

	switch {
	case isComment(current):
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		current = scanner.Bytes()
		fallthrough
	case isElement(current):
		doc.Root = parseElement(current)
		doc.Root.addNamespaces()
		if doc.Root.Closed {
			return doc, nil
		}

		err := scanTree(scanner, doc.Root, []*Element{doc.Root})
		if err != nil {
			return nil, err
		}

		return doc, err
	case isText(current):
		return nil, fmt.Errorf("Garbage found where root element should be")
	default:
		return nil, fmt.Errorf("No root element found")
	}

	return doc, nil
}

// Internal -------------------------------------------------------------------

func scanTree(scanner *bufio.Scanner, parent *Element, stack []*Element) error {
	for {
		scanner.Scan()
		current := scanner.Bytes()
		if err := scanner.Err(); err != nil {
			return err
		}
		//fmt.Printf("depth: %v, parent: %#v, token: %#v",
		//	len(stack), parent.Name.Name, string(current))

		switch {
		case isComment(current):
			continue
		case isEndElement(current):
			//fmt.Println(", type: end")
			if len(stack) == 1 {
				return nil
			}
			l := len(stack)
			stack = stack[:l-1]
			parent = stack[l-2]
			continue
		case isElement(current):
			//fmt.Println(", type: elem")
			el := parseElement(current)
			el.Parent = parent
			parent.Children = append(parent.Children, el)
			el.addNamespaces()
			if el.Closed {
				continue
			}
			parent, stack = el, append(stack, el)
			continue
		case isText(current):
			//fmt.Println(", type: text")
			parent.Text = string(unescapeText(current))
			continue
		}
	}

	return nil
}

func parseElement(data []byte) *Element {
	e := NewElement()
	if isClosed(data) {
		e.Closed = true
	}

	if hasAttribute(data) {
		ws := bytes.Index(data, []byte{' '})
		e.Name = parseQName(data[1:ws])
		attrs := parseAttributes(data[ws+1:])
		e.Attributes = attrs
	} else {
		if e.Closed {
			e.Name = parseQName(data[1 : len(data)-2])
		} else {
			e.Name = parseQName(data[1 : len(data)-1])
		}
	}

	return e
}

type attribute struct {
	name  QName
	value string
}

func parseAttributes(data []byte) []attribute {
	as := []attribute{}

	for {
		eq := bytes.Index(data, []byte{'='})
		if eq == -1 {
			break
		}
		qname := parseQName(data[:eq])

		endQuote := data[eq+1]
		end := bytes.Index(data[eq+2:], []byte{endQuote})
		val := string(unescapeText(data[eq+2 : eq+2+end]))
		as = append(as, attribute{
			name:  qname,
			value: val,
		})
		data = data[eq+end+4:]
	}

	return as
}
