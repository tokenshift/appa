package appa

import (
	"strings"
	"testing"
)

func Test_LiteralParseString(t *testing.T) {
	input := CreateStringBuffer(strings.NewReader("foo"))

	g := NewGrammar()

	result, err := g.Lit("foo").Parse(input)
	if err != nil {
		t.Error(err)
		return
	}

	assertStringerEquals(t, "foo", result[0].Val())
	assertIntEquals(t, 0, len(result[0].Children()))
}
