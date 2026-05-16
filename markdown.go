package markdowntohtml

import (
	"bufio"
	"fmt"
	"strings"
)

type Tag int
const (
	heading1 Tag = iota
	heading2
	heading3
	heading4
	heading5
	heading6
)

type MdConverter struct {
	Stack []Tag
	Source bufio.Scanner
	Result string
}

func (c MdConverter) Convert() error {
	for c.Source.Scan() {
		line := c.Source.Text()
		fmt.Println(line)
	}
	return nil
}

func Md2html(input string) (string, error) {
	con := MdConverter {
		Source: *bufio.NewScanner(strings.NewReader(input)),
	}
	
	err := con.Convert()
		
	return con.Result, err
}