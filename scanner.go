package xml

import (
	"bufio"
	"bytes"
	"fmt"
	"unicode"
)

var (
	tagStart = []byte{'<'}
	tagEnd   = []byte{'>'}
)

func split(data []byte, atEOF bool) (int, []byte, error) {
	tmpData := bytes.TrimLeft(data, " \n\r\t")
	diff := len(data) - len(tmpData)
	data = tmpData

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	} else if atEOF && !isWhitespace(data) {
		return len(data), data, bufio.ErrFinalToken
	}

	switch {
	case bytes.HasPrefix(data, tagStart):
		end := bytes.Index(data, tagEnd)
		if end == -1 {
			return 0, nil, nil
		}
		token := data[:end+1]

		next := bytes.Index(data, tagStart)
		if next == -1 {
			return 0, nil, nil
		} else if isWhitespace(data[end+1 : end+1+next]) {
			return len(token) + diff + next, token, nil
		}

		return len(token) + diff, token, nil
	default:
		end := bytes.Index(data, tagStart)
		if end == -1 {
			return 0, nil, nil
		}
		token := data[:end]
		return len(token) + diff, bytes.TrimSpace(token), nil
	}

	return 0, nil, fmt.Errorf("Not start of tag")
}

func isWhitespace(data []byte) bool {
	for _, c := range bytes.Runes(data) {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

func checkIfLast(data []byte) (int, error) {
	ind := bytes.Index(data, []byte{'<'})
	if ind == -1 {
		return ind, bufio.ErrFinalToken
	}
	return ind, nil
}
