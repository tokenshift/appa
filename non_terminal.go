package appa

type NonTerminal interface {
	Token

	Match(tokens ...Token) Rule
}

type nonTerminal struct {
}
