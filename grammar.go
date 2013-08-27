package appa

import "fmt"
import "io"
import "regexp"

type Grammar interface {
	Compile(start NonTerminal) (p Parser, err error)

	Lit(text string) Terminal
	Regex(pattern string) Terminal

	NonTerm(name string) NonTerminal

	WriteLALRCollection(start NonTerminal, out io.Writer)
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
	s, ok := start.(*nonTerminal)
	if !ok {
		panic(fmt.Sprintf("%v is not a non-terminal in this grammar.", start))
	}

	lexer := createLexer(g)
	collection := createLALRCollection(s)
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
			make([]rule, 0),
			"",
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

// Writes the collection of LALR sets in DOT format to
// the specified output stream. Useful for debugging
// parser construction.
func (g *grammar) WriteLALRCollection(start NonTerminal, out io.Writer) {
	s, ok := start.(*nonTerminal)
	if !ok {
		panic(fmt.Sprintf("%v is not a non-terminal in this grammar.", start))
	}

	collection := createLALRCollection(s)
	collection.writeTo(out)
}
