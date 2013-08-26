package appa

import "fmt"
import "regexp"

type Grammar interface {
	Compile(start NonTerminal) (p Parser, err error)

	Lit(text string) Terminal
	Regex(pattern string) Terminal

	NonTerm(name string) NonTerminal
}

type grammar struct {
	literals map[string]lit
	regexes map[string]regex

	nonterminals map[string]nonTerminal
}

func NewGrammar() Grammar {
	g := new(grammar)

	g.literals = make(map[string]lit)
	g.regexes = make(map[string]regex)

	g.nonterminals = make(map[string]nonTerminal)

	return g
}

// Compiles the grammar into a parser, with the specified
// non-terminal as the start symbol.
func (g *grammar) Compile(start NonTerminal) (p Parser, err error) {
	lexer := createLexer(g)
	collection := createLALRCollection(g)
	states := collection.createTable();
	p = parser { lexer, states }

	return nil, fmt.Errorf("Grammar.Compile not implemented.")
}

// Creates a token that will match the exact text specified.
func (g *grammar) Lit(text string) Terminal {
	var l lit
	var ok bool
	if l, ok = g.literals[text]; !ok {
		l = lit(text)
		g.literals[text] = l
	}

	return l
}

// Creates or retrieves a non-terminal with the specified name.
func (g *grammar) NonTerm(name string) NonTerminal {
	var nt nonTerminal
	var ok bool
	if nt, ok = g.nonterminals[name]; !ok {
		nt = nonTerminal {
			name,
			make([][]Token, 0),
		}
		g.nonterminals[name] = nt
	}

	return &nt
}

// Creates a token that will match the specified regular expression.
func (g *grammar) Regex(pattern string) Terminal {
	var r regex
	var ok bool
	if r, ok = g.regexes[pattern]; !ok {
		r = regex {
			regexp.MustCompile(pattern),
		}
		g.regexes[pattern] = r
	}

	return r
}
