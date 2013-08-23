package appa

// A single production rule for the grammar.
type production struct {
	nt *nonTerminal
	body []Token
}
