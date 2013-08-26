package appa

import "fmt"

type Parser interface {
	ParseString(input string) (result Node, err error)
}

type parser struct {
	lexer lexer
	states actionTable
}

func (p parser) ParseString(input string) (result Node, err error) {
	return nil, fmt.Errorf("ParseString not implemented.")
}
