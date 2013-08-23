package appa

import "fmt"
import "testing"

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
