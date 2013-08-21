package appa

import "fmt"

// A string literal that can act as
// either a parse rule or a node.
type Lit string

func (lit Lit) And(r Rule) Rule {
	rules := make([]Rule, 2, 2)
	rules[0] = lit
	rules[1] = r
	return sequence {
		rules,
	}
}

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

func (lit Lit) Parse(input StringBuffer) (ast Node, err error) {
	if matched := lit.Match(input, 0); matched > 0 {
		input.Discard(matched)
		ast = lit
	} else {
		ast = nil
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
