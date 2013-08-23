package appa

import "fmt"

type Grammar struct {
}

// Compiles the grammar into a parser, with the specified
// non-terminal as the start symbol.
func (g *Grammar) Compile(start NonTerminal) (p Parser, err error) {
	return nil, fmt.Errorf("Grammar.Compile not implemented.")
}

// Creates a token that will match the exact text specified.
func (g *Grammar) Lit(text string) Terminal {
	return nil
}

// Creates or retrieves a non-terminal with the specified name.
func (g *Grammar) NonTerm(name string) NonTerminal {
	return nil
}

// Creates a token that will match the specified regular expression.
func (g *Grammar) Regex(pattern string) Terminal {
	return nil
}
