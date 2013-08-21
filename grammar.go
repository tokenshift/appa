package appa

import "regexp"

type Grammar interface {
	Lit(value string) Lit
	NonTerminal(name string) NonTerminal
	Regex(pattern string) Regex
}

func NewGrammar() Grammar {
	g := new(grammar)

	g.nonterminals = make(map[string]NonTerminal)
	g.rules = make(map[string][]Rule)

	return g
}

type grammar struct {
	nonterminals map[string]NonTerminal
	rules map[string][]Rule
}

func (g *grammar) Lit(text string) (lit Lit) {
	return Lit(text)
}

func (g *grammar) Regex(pattern string) Regex {
	rx := regexp.MustCompile(pattern)
	return Regex { rx }
}
