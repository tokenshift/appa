package appa

import "fmt"

type NonTerminal interface {
	Token

	Match(tokens ...Token) Rule
}

type Rule interface {
}

type nonTerminal struct {
	name string
	rules [][]Token

	// Tag for system-generated non-terminals.
	special string
}

func (nt *nonTerminal) Match(tokens ...Token) Rule {
	nt.rules = append(nt.rules, tokens)

	return nil
}

func (nt *nonTerminal) String() string {
	if nt.special == "" {
		return fmt.Sprintf("<%s>", nt.name)
	} else {
		return nt.special
	}
}
