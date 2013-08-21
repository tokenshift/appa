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

	exp.AddRule(Seq(num, oper, exp))
	exp.AddRule(num)

	result, err := exp.Parse(source)
	if err != nil {
		t.Error(err)
		return
	}

	assertStringerEquals(t, "EXP", result[0].Val())
	assertIntEquals(t, 3, len(result[0].Children()))
	assertNodeStringEquals(t, "(EXP 1 + (EXP 2 + (EXP 3)))", result[0])
}
