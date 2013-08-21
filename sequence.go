package appa

import (
	"bytes"
	"fmt"
)

// Concatenates the rules into a single rule.
func Seq(rules ...Rule) Rule {
	return Sequence(rules)
}

type Sequence []Rule

func (s Sequence) Match(input StringBuffer, offset int) int {
	origOffset := offset

	for _, rule := range(s) {
		matched := rule.Match(input, offset)
		if matched < 0 {
			return -1
		}

		offset += matched
	}

	return offset - origOffset
}

func (s Sequence) Parse(input StringBuffer) (node Node, err error) {
	if s.Match(input, 0) <= 0 {
		node = nil
		err = fmt.Errorf("Failed to match sequence: %s", s)
		return
	}

	nodes := make([]Node, len(s), len(s))
	for i, rule := range(s) {
		node, err = rule.Parse(input)
		if err != nil {
			return
		}

		nodes[i] = node
	}

	node = NodeList(nodes)
	return
}

func (s Sequence) String() string {
	var buffer bytes.Buffer

	first := true

	for _, rule := range(s) {
		if !first {
			buffer.WriteString(" ")
		}

		buffer.WriteString(fmt.Sprint(rule))

		first = false
	}

	return buffer.String()
}
