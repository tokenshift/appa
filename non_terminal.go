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
	rules []rule

	// Tag for system-generated non-terminals.
	special string
}

func (nt *nonTerminal) Equals(other Token) bool {
	if nt2, ok := other.(*nonTerminal); ok {
		return nt2 == nt
	} else {
		return false
	}
}

func (nt *nonTerminal) Match(tokens ...Token) Rule {
	nt.rules = append(nt.rules, createRule(tokens...))

	return nil
}

func (nt *nonTerminal) String() string {
	if nt.special == "" {
		return fmt.Sprintf("<%s>", nt.name)
	} else {
		return nt.special
	}
}
