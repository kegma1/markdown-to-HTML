package markdowntohtml_test

import (
	"fmt"
	"testing"
	"github.com/kegma1/markdown-to-HTML"
)

func TestH1(t *testing.T) {
	msg := "Hallo verden"
	
	html, err := markdowntohtml.Md2html(fmt.Sprintf("# %s", msg))
	
	if err != nil {
		t.Error(err)
		return
	}

	if html != fmt.Sprintf("<h1>%s</h1>", msg) {
		t.Log(html)
		t.Fail()
	}
}

func TestMultiline(t *testing.T) {
	msg := "Hallo\nVerden\ner\nmin"
	markdowntohtml.Md2html(msg)
}