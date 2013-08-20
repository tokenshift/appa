package appa

import (
	"bytes"
	"fmt"
)

type sequence struct {
	rules []Rule
}

func makeSequence(rules ...Rule) sequence {
	var seq sequence
	seq.rules = make([]Rule, len(rules), len(rules))

	for i, rule := range(rules) {
		seq.rules[i] = rule
	}

	return seq
}

func (s sequence) And(r Rule) Rule {
	return &sequence {
		append(s.rules, r),
	}
}

func (s sequence) Match(input StringBuffer, offset int) int {
	origOffset := offset

	for _, rule := range(s.rules) {
		matched := rule.Match(input, offset)
		if matched < 0 {
			return -1
		}

		offset += matched
	}

	return offset - origOffset
}

func (s sequence) Parse(input StringBuffer) (ast Node, err error) {
	if s.Match(input, 0) <= 0 {
		err = fmt.Errorf("Failed to match sequence: %s", s)
		return
	}

	nodes := make([]Node, len(s.rules), len(s.rules))
	for i, rule := range(s.rules) {
		ast, err = rule.Parse(input)
		if err != nil {
			return
		}

		nodes[i] = ast
	}

	ast = Node {
		"",
		nodes,
	}

	return
}

func (s sequence) String() string {
	var buffer bytes.Buffer

	first := true

	for _, rule := range(s.rules) {
		if !first {
			buffer.WriteString(" ")
		}

		buffer.WriteString(fmt.Sprint(rule))

		first = false
	}

	return buffer.String()
}

// Trim applies to all of the children of the sequence.
func (s sequence) Trim() Rule {
	seq := sequence {
		make([]Rule, len(s.rules), len(s.rules)),
	}

	for i, rule := range(s.rules) {
		seq.rules[i] = rule.Trim()
	}
	
	return seq
}
