package appa

import (
	"bytes"
	"fmt"
)

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

func (s Sequence) Parse(input StringBuffer) (result []Node, err error) {
	if s.Match(input, 0) <= 0 {
		err = fmt.Errorf("Failed to match sequence: %s", s)
		return
	}

	nodes := make([]Node, len(s), len(s))
	for i, rule := range(s) {
		result, err = rule.Parse(input)
		if err != nil {
			return
		}

		if result == nil || len(result) == 0 {
			nodes[i] = nil
		} else if len(result) == 1 {
			nodes[i] = result[0]
		} else {
			nodes[i] = NodeList(result)
		}
	}

	result = nodes
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
