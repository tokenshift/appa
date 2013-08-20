package appa

import (
	"strings"
	"testing"
)

func Test_OptionalLiteral(t *testing.T) {
	g := NewGrammar()

	nt := g.NonTerminal("NT")
	nt.AddRule(g.Literal("111").
		And(g.Literal("222")).
		And(g.Literal("333").Optional()).
		And(g.Literal("444")).
		And(g.Literal("555")))


	input := CreateStringBuffer(strings.NewReader("111222444555"))
	ast, err := nt.Parse(input)

	if err != nil {
		t.Error(err)
	}

	assertStringEquals(t, "(NT 111 222 () 444 555)", ast.String())


	input = CreateStringBuffer(strings.NewReader("111222333444555"))
	ast, err = nt.Parse(input)

	if err != nil {
		t.Error(err)
	}

	assertStringEquals(t, "(NT 111 222 333 444 555)", ast.String())
}
