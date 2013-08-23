package appa

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
