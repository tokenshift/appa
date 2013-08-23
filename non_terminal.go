package appa

import "fmt"

type NonTerminal interface {
	Token

	Match(tokens ...Token) Rule
}

type nonTerminal struct {
	name string
}

func (nt *nonTerminal) Match(tokens ...Token) Rule {
	return nil
}

func (nt *nonTerminal) String() string {
	return fmt.Sprintf("<%s>", nt.name)
}
