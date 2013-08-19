package appa

import (
	"fmt"
)

type Literal struct {
	Text string
}

func (g *grammar) Literal(text string) (lit Literal) {
	lit, ok := g.literals[text]
	if !ok {
		lit = Literal {
			text,
		}
	}
	return
}

func (lit Literal) And(r Rule) Rule {
	rules := make([]Rule, 2, 2)
	rules[0] = lit
	rules[1] = r
	return sequence {
		rules,
	}
}

func (lit Literal) ParseString(input string) (ast Node, err error) {
	if input == lit.Text {
		ast = Node {
			input,
			make([]Node, 0, 0),
		}
	} else {
		err = fmt.Errorf("Expected literal '%v'.", lit.Text)
	}

	return
}
