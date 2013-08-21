package appa

type grammar struct {
	nonterminals map[string]NonTerminal
	rules map[string][]Rule
}

func (g *grammar) Lit(text string) Rule {
	return Lit(text)
}

func (g *grammar) NonTerminal(name string) (nt NonTerminal) {
	nt, ok := g.nonterminals[name]
	if !ok {
		nt = nonTerminal {
			g,
			name,
		}
		g.nonterminals[name] = nt
		g.rules[name] = make([]Rule, 0, 0)
	}
	return
}

func (g *grammar) Regex(pattern string) Rule {
	return Regex(pattern)
}
