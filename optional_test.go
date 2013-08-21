package appa

import (
	"strings"
	"testing"
)

func Test_OptionalLit(t *testing.T) {
	g := NewGrammar()

	nt := g.NonTerminal("NT")
	nt.AddRule(Seq(
		g.Lit("111"),
		g.Lit("222"),
		Opt(g.Lit("333")),
		g.Lit("444"),
		g.Lit("555")))


	input := CreateStringBuffer(strings.NewReader("111222444555"))
	ast, err := nt.Parse(input)

	if err != nil {
		t.Error(err)
	}

	assertNodeStringEquals(t, "(NT 111 222 () 444 555)", ast)


	input = CreateStringBuffer(strings.NewReader("111222333444555"))
	ast, err = nt.Parse(input)

	if err != nil {
		t.Error(err)
	}

	assertNodeStringEquals(t, "(NT 111 222 333 444 555)", ast)
}
