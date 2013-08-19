package appa

import (
	"strings"
	"testing"
)

func Test_ReadNonTerminal(t *testing.T) {
	g := NewGrammar()

	foo := g.NonTerminal("FOO")

	foo.AddRule(g.Literal("foo"))
	foo.AddRule(g.Literal("bar"))

	input := CreateStringBuffer(strings.NewReader("testfoobartest"))

	if foo.Match(input, 4) != 3 {
		t.Errorf("Expected \"foo\" at position 4")
	}

	if foo.Match(input, 7) != 3 {
		t.Errorf("Expected \"bar\" at position 7")
	}

	if foo.Match(input, 0) != -1 {
		t.Errorf("Did not expect a match at position 0")
	}

	input = CreateStringBuffer(strings.NewReader("foobar"))

	lit1, err := foo.Parse(input)
	if err != nil {
		t.Error(err)
	}

	lit2, err := foo.Parse(input)
	if err != nil {
		t.Error(err)
	}

	assertStringEquals(t, "(FOO foo)", lit1.String())
	assertStringEquals(t, "(FOO bar)", lit2.String())
}
