package parser

import (
	"fmt"
	"strings"
)

const TAB string = "\t"

var indentDepth int

func insertIndent() string {
	return strings.Repeat(TAB, indentDepth)
}

func trace(msg string) string {
	fmt.Printf("%s%s\n", insertIndent(), "BEG "+msg)

	indentDepth++

	return msg
}

func untrace(msg string) {
	indentDepth--

	fmt.Printf("%s%s\n", insertIndent(), "END "+msg)
}
