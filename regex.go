package appa

import (
	"fmt"
	"regexp"
)

type Regex struct {
	Pattern *regexp.Regexp
}

func (rx Regex) And(r Rule) Rule {
	rules := make([]Rule, 2, 2)
	rules[0] = rx
	rules[1] = r
	return sequence {
		rules,
	}
}

func (rx Regex) Match(input StringBuffer, offset int) int {
	if ok, match := input.ReadPattern(rx.Pattern, offset); ok {
		return len(match)
	} else {
		return -1
	}
}

func (rx Regex) Parse(input StringBuffer) (node Node, err error) {
	if matches := rx.Match(input, 0); matches > 0 {
		text := input.Consume(matches)
		node = Lit(text)
	} else {
		node = nil
		err = fmt.Errorf("Input string did not match pattern {0}", rx.Pattern)
	}

	return
}

func (rx Regex) String() string {
	return fmt.Sprintf("/%v/", rx.Pattern)
}
