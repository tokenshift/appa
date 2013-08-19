package appa

import (
	"fmt"
)

type sequence struct {
	rules []Rule
}

func (s sequence) And(r Rule) Rule {
	return &sequence {
		append(s.rules, r),
	}
}

func (s sequence) ParseString(input string) (ast Node, err error) {
	err = fmt.Errorf("Not implemented")
	return
}
