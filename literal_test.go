package appa

import (
	"strings"
	"testing"
)

func Test_LiteralParseString(t *testing.T) {
	input := CreateStringBuffer(strings.NewReader("foo"))

	g := NewGrammar()

	ast, err := g.Lit("foo").Parse(input)
	if err != nil {
		t.Error(err)
		return
	}

	assertStringerEquals(t, "foo", ast.Val())
	assertIntEquals(t, 0, len(ast.Children()))
}
