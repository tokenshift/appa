package appa

import (
	"strings"
	"testing"
)

func Test_RegexParseString(t *testing.T) {
	input := CreateStringBuffer(strings.NewReader("12345"))

	g := NewGrammar()

	rx := g.Regex("\\d+")

	ast, err := rx.Parse(input)
	if err != nil {
		t.Error(err)
		return
	}

	assertStringerEquals(t, "12345", ast.Val())
	assertIntEquals(t, 0, len(ast.Children()))

	input = CreateStringBuffer(strings.NewReader("foofoofoo"))

	rx = g.Regex("[a-z]{3}")

	assertIntEquals(t, 3, rx.Match(input, 0))
	assertIntEquals(t, 3, rx.Match(input, 3))
	assertIntEquals(t, 3, rx.Match(input, 6))
}
