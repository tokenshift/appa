package appa

import (
	"fmt"
	"regexp"
)

type Regex struct {
	Pattern *regexp.Regexp
}

func (g *grammar) Regex(pattern string) (rx Regex, err error) {
	rx, ok := g.regexes[pattern]
	if !ok {
		var regex *regexp.Regexp
		regex, err = regexp.Compile(pattern)
		if err != nil {
			return
		}

		rx = Regex {
			regex,
		}
		g.regexes[pattern] = rx
	}
	return
}

func (rx Regex) And(r Rule) Rule {
	rules := make([]Rule, 2, 2)
	rules[0] = rx
	rules[1] = r
	return sequence {
		rules,
	}
}

func (rx Regex) ParseString(input string) (ast Node, err error) {
	match := rx.Pattern.FindStringIndex(input)
	if match == nil {
		err = fmt.Errorf("Input string did not match regular expression.")
		return
	}

	ast = Node {
		input[match[0]:match[1]],
		make([]Node, 0, 0),
	}

	return
}
