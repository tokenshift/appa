package appa

import (
	"strings"
	"testing"
)

func Test_RegexParseString(t *testing.T) {
	input := CreateStringBuffer(strings.NewReader("12345"))

	g := NewGrammar()

	rx, err := g.Regex("\\d+")
	if err != nil {
		t.Error(err)
		return
	}

	ast, err := rx.Parse(input)
	if err != nil {
		t.Error(err)
		return
	}

	if ast.Name != "12345" {
		t.Errorf("Expected \"%v\", got \"%v\".", "12345", ast.Name)
	}
	if len(ast.Children) > 0 {
		t.Errorf("Expected %d children, got %d.", len(ast.Children))
	}
}
