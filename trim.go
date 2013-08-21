package appa

import (
	"fmt"
	"regexp"
)

var rx_trim *regexp.Regexp = regexp.MustCompile("\\s+")

type trim struct {
	rule Rule
}

func (t trim) Match(buffer StringBuffer, offset int) int {
	ok, match := buffer.ReadPattern(rx_trim, offset)
	if ok {
		offset += len(match)
	}
	
	matched := t.rule.Match(buffer, offset)
	if matched > -1 {
		return matched + len(match)
	}

	return -1
}

func (t trim) Parse(buffer StringBuffer) (ast Node, err error) {
	ok, match := buffer.ReadPattern(rx_trim, 0)
	if ok {
		buffer.Discard(len(match))
	}

	ast, err = t.rule.Parse(buffer)
	return
}

func (t trim) String() string {
	return fmt.Sprintf("{trim} %v", t.rule)
}
