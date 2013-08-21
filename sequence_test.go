package appa

import (
	"strings"
	"testing"
)

func Test_SimpleSequence(t *testing.T) {
	g := NewGrammar()

	seq := Sequence([]Rule {
		g.Lit("1"),
		g.Lit("2"),
		g.Lit("3"),
	})

	input := CreateStringBuffer(strings.NewReader("123"))

	assertIntEquals(t, 3, seq.Match(input, 0))
}
