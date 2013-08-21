package appa

import (
	"strings"
	"testing"
)

func Test_SimpleProgram(t *testing.T) {
	var source = CreateStringBuffer(strings.NewReader("1+2+3"))

	var g = NewGrammar()

	exp := g.NonTerminal("EXP")

	num := g.Regex("\\d+")
	oper := g.Lit("+")

	exp.AddRule(num.And(oper).And(exp))
	exp.AddRule(num)

	ast, err := exp.Parse(source)
	if err != nil {
		t.Error(err)
		return
	}

	assertStringerEquals(t, "EXP", ast.Val())
	assertIntEquals(t, 3, len(ast.Children()))
	assertNodeStringEquals(t, "(EXP 1 + (EXP 2 + (EXP 3)))", ast)
}
