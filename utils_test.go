package xml

import (
	"bytes"
	"testing"
)

func TestPredicates(t *testing.T) {
	datas := [][]byte{
		[]byte("<?xml version=\"1.0\" encoding=\"UTF-8\" ?>"),
		[]byte("<abc/>"),
		[]byte("<!-- comment! åöä€ -->"),
		[]byte("<he:llo stu:ff=\"tns:val\">"),
		[]byte("</he:llo>"),
	}

	switch {
	case !isProlog(datas[0]):
		t.Error("Failed to recognize prolog")
	case !isClosed(datas[1]):
		t.Error("Failed to recognize closed")
	case !isComment(datas[2]):
		t.Error("Failed to recognize comment")
	case !isElement(datas[3]):
		t.Error("Failed to recognize element")
	case !isElement(datas[3]):
		t.Error("Failed to recognize element")
	case !isEndElement(datas[4]):
		t.Error("Failed to recognize end-element")
	}
}

func TestEscapeText(t *testing.T) {
	text := []byte(`15 < 23 && len("str") > 1 & also 'this'`)
	expects := []byte(`15 &lt; 23 &amp;&amp; len(&quot;str&quot;) &gt; 1 &amp; also &apos;this&apos;`)

	if esc := escapeText(text); bytes.Compare(esc, expects) != 0 {
		t.Error("Failed to escaping text")
	}
}

func TestUnescapeText(t *testing.T) {
	esc := []byte(`15 &lt; 23 &amp;&amp; len(&quot;str&quot;) &gt; 1 &amp; also &apos;this&apos;`)
	expects := []byte(`15 < 23 && len("str") > 1 & also 'this'`)

	if text := unescapeText(esc); bytes.Compare(text, expects) != 0 {
		t.Error("Failed to unescaping text")
	}
}
