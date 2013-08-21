package appa

import "fmt"

func (lit Lit) Children() []Node {
	return Empty
}

func (lit Lit) Match(input StringBuffer, offset int) int {
	if input.ReadLiteral(string(lit), offset) {
		return len(lit)
	} else {
		return -1
	}
}

func (lit Lit) Parse(input StringBuffer) (result []Node, err error) {
	if matched := lit.Match(input, 0); matched > 0 {
		input.Discard(matched)
		result = []Node{lit}
	} else {
		err = fmt.Errorf("Expected literal '%v'.", lit)
	}

	return
}

func (lit Lit) String() string {
	return string(lit)
}

func (lit Lit) Val() fmt.Stringer {
	return lit
}
