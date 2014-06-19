package main

import "github.com/tokenshift/appa"

func main() {
	g := appa.CreateGrammar()

	op_add := g.Literal("+")
	lit_num := g.Regex("\\d+")

	expr := g.NonTerminal("expr")
	expr.AddRule(expr, op_add, lit_num)
	expr.AddRule(lit_num)

	g.Compile(expr)
}
