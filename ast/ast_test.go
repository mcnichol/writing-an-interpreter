package ast

import (
	"../token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token:token.Token{Type:token.IDENT, Literal: "anotherVar"},
					Value:"anotherVar",
				},
			},
		},
	}

	expectedProgram := "let myVar = anotherVar;"

	if program.String() != expectedProgram {
		t.Errorf("\nprogram.String() wrong.\nexpected:\t%q\nactual:\t\t%q", expectedProgram, program.String())
	}
}
