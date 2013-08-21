package appa

import (
	"strings"
	"testing"
)

func Test_SimpleSequence(t *testing.T) {
	g := NewGrammar()

	seq := sequence {
		make([]Rule, 0, 0),
	}

	seq.rules = append(seq.rules, g.Lit("1"))
	seq.rules = append(seq.rules, g.Lit("2"))
	seq.rules = append(seq.rules, g.Lit("3"))

	input := CreateStringBuffer(strings.NewReader("123"))

	assertIntEquals(t, 3, seq.Match(input, 0))
}
