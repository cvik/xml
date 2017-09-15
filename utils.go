package xml

import "bytes"

func isProlog(data []byte) bool {
	return predicate(data, "<?xml", "?>")
}

func isElement(data []byte) bool {
	return predicate(data, "<", ">")
}

func isClosed(data []byte) bool {
	return predicate(data, "", "/>")
}

func isEndElement(data []byte) bool {
	return predicate(data, "</", "")
}

func isComment(data []byte) bool {
	return predicate(data, "<!--", "-->")
}

func isText(data []byte) bool {
	return !bytes.ContainsAny(data, "<>\"")
}

func hasAttribute(data []byte) bool {
	return bytes.Contains(data, []byte{'='})
}

func predicate(data []byte, prefix, suffix string) bool {
	hasPrefix := bytes.HasPrefix(data, []byte(prefix))
	hasSuffix := bytes.HasSuffix(data, []byte(suffix))
	if hasPrefix && hasSuffix {
		return true
	}
	return false
}

func escapeText(data []byte) []byte {
	data = bytes.Replace(data, []byte(`&`), []byte(`&amp;`), -1)
	data = bytes.Replace(data, []byte(`"`), []byte(`&quot;`), -1)
	data = bytes.Replace(data, []byte(`'`), []byte(`&apos;`), -1)
	data = bytes.Replace(data, []byte(`<`), []byte(`&lt;`), -1)
	data = bytes.Replace(data, []byte(`>`), []byte(`&gt;`), -1)

	return data
}

func unescapeText(data []byte) []byte {
	data = bytes.Replace(data, []byte(`&quot;`), []byte(`"`), -1)
	data = bytes.Replace(data, []byte(`&apos;`), []byte(`'`), -1)
	data = bytes.Replace(data, []byte(`&lt;`), []byte(`<`), -1)
	data = bytes.Replace(data, []byte(`&gt;`), []byte(`>`), -1)
	data = bytes.Replace(data, []byte(`&amp;`), []byte(`&`), -1)

	return data
}
