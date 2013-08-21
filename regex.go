package appa

import (
	"fmt"
	"regexp"
)

type regex struct {
	pattern *regexp.Regexp
}

func (rx regex) Match(input StringBuffer, offset int) int {
	if ok, match := input.ReadPattern(rx.pattern, offset); ok {
		return len(match)
	} else {
		return -1
	}
}

func (rx regex) Parse(input StringBuffer) (node Node, err error) {
	if matches := rx.Match(input, 0); matches > 0 {
		text := input.Consume(matches)
		node = Lit(text)
	} else {
		node = nil
		err = fmt.Errorf("Input string did not match pattern {0}", rx.pattern)
	}

	return
}

func (rx regex) String() string {
	return fmt.Sprintf("/%v/", rx.pattern)
}
