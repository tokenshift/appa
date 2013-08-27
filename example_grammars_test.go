package appa

import "fmt"
import "testing"

func createExpressionGrammar() (g Grammar, start NonTerminal) {
	g = NewGrammar()

	e := g.NonTerm("E")
	t := g.NonTerm("T")
	v := g.NonTerm("V")

	op1 := g.NonTerm("OP1")
	op1.Match(g.Lit("+"))
	op1.Match(g.Lit("-"))

	op2 := g.NonTerm("OP2")
	op2.Match(g.Lit("*"))
	op2.Match(g.Lit("/"))

	// E -> E + T | T
	e.Match(e, op1, t)
	e.Match(t)

	// T -> T * V | V
	t.Match(t, op2, v)
	t.Match(v)

	// V -> ( E ) | n
	v.Match(g.Lit("("), e, g.Lit(")"))
	v.Match(g.Regex("\\d+"))

	return g, e
}

func Test_SimpleExpressionGrammar(t *testing.T) {
	g := NewGrammar()

	expr := g.NonTerm("E")
	term := g.NonTerm("T")
	fact := g.NonTerm("F")

	// E -> E + T | T
	expr.Match(expr, g.Lit("+"), term)
	expr.Match(term)

	// T -> T * F | F
	term.Match(term, g.Lit("*"), fact)
	term.Match(fact)

	// F -> ( E ) | n
	fact.Match(g.Lit("("), expr, g.Lit(")"))
	fact.Match(g.Regex("\\d+"))

	parser, err := g.Compile(expr)

	if err != nil {
		t.Error(err)
		return
	}

	result, err := parser.ParseString("17+3*(8+2)+4")

	if err != nil {
		t.Error(err)
		return
	}

	// (E (E 17 + (T 3 * (F (E 8 + 2)))) + 4)
	assertStringEquals(t, "(E (E 17 + (T 3 * (F (E 8 + 2)))) + 4)", fmt.Sprint(result))
}
