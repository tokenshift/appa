package appa

type Grammar interface {
	Literal(value string) (lit Literal)
	NonTerminal(name string) (nt NonTerminal)
	Regex(pattern string) (rx Regex, err error)
}

func NewGrammar() Grammar {
	g := new(grammar)

	g.literals = make(map[string]Literal)
	g.nonterminals = make(map[string]NonTerminal)
	g.regexes = make(map[string]Regex)

	return g
}

type Rule interface {
	And(r Rule) Rule
	ParseString(input string) (Node, error)
}

type grammar struct {
	literals map[string]Literal
	nonterminals map[string]NonTerminal
	regexes map[string]Regex
}
