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

func (lit Literal) Match(input StringBuffer, offset int) int {
	if input.ReadLiteral(lit.Text, offset) {
		return len(lit.Text)
	} else {
		return -1
	}
}

func (lit Literal) Parse(input StringBuffer) (ast Node, err error) {
	if matched := lit.Match(input, 0); matched > 0 {
		input.Discard(matched)
		ast = Node {
			lit.Text,
			make([]Node, 0, 0),
		}
	} else {
		err = fmt.Errorf("Expected literal '%v'.", lit.Text)
	}

	return
}
