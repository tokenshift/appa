package appa

import "fmt"

type optional struct {
	rule Rule
}

func (o optional) Match(buffer StringBuffer, offset int) int {
	matched := o.rule.Match(buffer, offset)
	
	if matched > 0 {
		return matched
	}

	return 0
}

func (o optional) Parse(buffer StringBuffer) (node []Node, err error) {
	matched := o.rule.Match(buffer, 0)
	if matched > 0 {
		node, err = o.rule.Parse(buffer)
	} else {
		node = []Node{}
	}

	return
}

func (o optional) String() string {
	return fmt.Sprintf("[%v]", o.rule)
}
