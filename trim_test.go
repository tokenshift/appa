package appa

import (
	"strings"
	"testing"
)

func Test_TrimLiteral(t *testing.T) {
	g := NewGrammar()
	foo := g.Literal("foo").Trim()

	ast, err := foo.Parse(CreateStringBuffer(strings.NewReader("foo")))

	if err != nil {
		t.Error(err)
	}

	assertStringEquals(t, "foo", ast.String())


	ast, err = foo.Parse(CreateStringBuffer(strings.NewReader(" \tfoo")))

	if err != nil {
		t.Error(err)
	}

	assertStringEquals(t, "foo", ast.String())
}

func Test_TrimSequence(t *testing.T) {
	g := NewGrammar()
	foo, _ := g.Regex("[a-z]{3}")
	test := g.NonTerminal("TEST")
	test.AddRule(foo.And(foo).And(foo).Trim())

	ast, err := test.Parse(CreateStringBuffer(strings.NewReader("foofoofoo")))

	if err != nil {
		t.Error(err)
	}

	assertStringEquals(t, "(TEST foo foo foo)", ast.String())


	ast, err = test.Parse(CreateStringBuffer(strings.NewReader(" foofoofoo")))

	if err != nil {
		t.Error(err)
	}
	

	ast, err = test.Parse(CreateStringBuffer(strings.NewReader("foofoo foo")))

	if err != nil {
		t.Error(err)
	}

	assertStringEquals(t, "(TEST foo foo foo)", ast.String())


	ast, err = test.Parse(CreateStringBuffer(strings.NewReader(" \tfoo foo\tfoo")))

	if err != nil {
		t.Error(err)
	}

	assertStringEquals(t, "(TEST foo foo foo)", ast.String())
}
