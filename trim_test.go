package appa

import (
	"strings"
	"testing"
)

func Test_TrimLit(t *testing.T) {
	g := NewGrammar()
	foo := Trim(g.Lit("foo"))

	ast, err := foo.Parse(CreateStringBuffer(strings.NewReader("foo")))

	if err != nil {
		t.Error(err)
	}

	assertStringerEquals(t, "foo", ast.Val())


	ast, err = foo.Parse(CreateStringBuffer(strings.NewReader(" \tfoo")))

	if err != nil {
		t.Error(err)
	}

	assertStringerEquals(t, "foo", ast.Val())
}

func Test_TrimSequence(t *testing.T) {
	g := NewGrammar()
	foo := g.Regex("[a-z]{3}")
	test := g.NonTerminal("TEST")
	test.AddRule(Trim(Seq(foo, foo, foo)))

	ast, err := test.Parse(CreateStringBuffer(strings.NewReader("foofoofoo")))

	if err != nil {
		t.Error(err)
	}

	assertNodeStringEquals(t, "(TEST foo foo foo)", ast)


	ast, err = test.Parse(CreateStringBuffer(strings.NewReader(" foofoofoo")))

	if err != nil {
		t.Error(err)
	}

	assertNodeStringEquals(t, "(TEST foo foo foo)", ast)
	

	ast, err = test.Parse(CreateStringBuffer(strings.NewReader("foofoo foo")))

	if err != nil {
		t.Error(err)
	}

	assertNodeStringEquals(t, "(TEST foo foo foo)", ast)


	ast, err = test.Parse(CreateStringBuffer(strings.NewReader(" \tfoo foo\tfoo")))

	if err != nil {
		t.Error(err)
	}

	assertNodeStringEquals(t, "(TEST foo foo foo)", ast)
}
