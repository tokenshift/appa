package appa

import (
	"testing"
)

func Test_LiteralParseString(t *testing.T) {
	input := "foo"

	g := NewGrammar()

	ast, err := g.Literal("foo").ParseString(input)
	if err != nil {
		t.Error(err)
		return
	}

	if ast.Name != "foo" {
		t.Errorf("Expected \"%v\", got \"%v\".", "foo", ast.Name)
	}
	if len(ast.Children) > 0 {
		t.Errorf("Expected %d children, got %d.", len(ast.Children))
	}
}
