package appa

import "fmt"
import "regexp"

type Grammar interface {
	NonTerminal(name string) NonTerminal

	Literal(val string) Terminal
	Regex(pattern string) Terminal
}

type Token interface {
	fmt.Stringer
}

func CreateGrammar() Grammar {
	return &grammar {
		make(map[string]*nonterm),
		make(map[string]lit),
		make(map[string]reg),
	}
}

type grammar struct {
	nts map[string]*nonterm
	lits map[string]lit
	regs map[string]reg
}

func (g *grammar) NonTerminal(name string) NonTerminal {
	if nt, ok := g.nts[name]; ok {
		return nt
	} else {
		nt = &nonterm {
			name,
			make([][]Token, 0),
		}
		g.nts[name] = nt
		return nt
	}
}

func (g *grammar) Literal(val string) Terminal {
	if t, ok := g.lits[val]; ok {
		return t
	} else {
		t = lit { val }
		g.lits[val] = t
		return t
	}
}

func (g *grammar) Regex(pattern string) Terminal {
	if t, ok := g.regs[pattern]; ok {
		return t
	} else {
		rx := regexp.MustCompile(pattern)
		t = reg { rx }
		g.regs[pattern] = t
		return t
	}
}
