package appa

import (
	"testing"
)

func Test_SimpleProgram(t *testing.T) {
	var source = "1+2+3"

	var g = NewGrammar()

	exp := g.NonTerminal("EXP")

	num, _ := g.Regex("\\d+")
	oper := g.Literal("+")

	exp.AddRule(num.And(oper).And(exp))
	exp.AddRule(num)

	ast, err := exp.ParseString(source)
	if err != nil {
		t.Error(err)
		return
	}

	if ast.Name != "EXP" {
		t.Errorf("Expected <EXP>")
	}

	if len(ast.Children) != 2 {
		t.Errorf("Expected 2 children")
	}

	expected := "(EXP 1 + (EXP 2 + (EXP 3)))"
	if ast.String() != expected {
		t.Errorf("Expected %v, got %v", expected, ast)
	}
}
