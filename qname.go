package xml

import "bytes"

type QName struct {
	Prefix string
	Name   string
}

func parseQName(data []byte) QName {
	if bytes.Contains(data, []byte(":")) {
		parts := bytes.SplitN(data, []byte(":"), 2)
		return QName{Prefix: string(parts[0]), Name: string(parts[1])}
	}
	return QName{Prefix: "", Name: string(data)}
}

func (n *QName) String() string {
	return n.Prefix + ":" + n.Name
}
