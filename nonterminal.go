package appa

import "fmt"

type NonTerminal interface {
	Token
	AddRule(tokens ...Token)
}

type nonterm struct {
	name string
	rules [][]Token
}

func (nt *nonterm) AddRule(tokens ...Token) {
	nt.rules = append(nt.rules, tokens)
}

func (nt *nonterm) String() string {
	return fmt.Sprintf("<%s>", nt.name)
}
