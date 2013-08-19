package appa

import (
	"fmt"
)

type NonTerminal struct {
	g *grammar
	Name string
}

func (g *grammar) NonTerminal(name string) (nt NonTerminal) {
	nt, ok := g.nonterminals[name]
	if !ok {
		nt = NonTerminal {
			g,
			name,
		}
		g.nonterminals[name] = nt
		g.rules[name] = make([]Rule, 0, 0)
	}
	return
}

func (nt NonTerminal) AddRule(r Rule) {
	nt.g.rules[nt.Name] = append(nt.g.rules[nt.Name], r)
}

func (nt NonTerminal) And(r Rule) Rule {
	rules := make([]Rule, 2, 2)
	rules[0] = nt
	rules[1] = r
	return sequence {
		rules,
	}
}

func (nt NonTerminal) Match(input StringBuffer, offset int) int {
	for _, rule := range nt.g.rules[nt.Name] {
		if matched := rule.Match(input, offset); matched > 0 {
			return matched
		}
	}

	return -1
}

func (nt NonTerminal) Parse(input StringBuffer) (ast Node, err error) {
	for _, rule := range nt.g.rules[nt.Name] {
		var result Node
		result, err = rule.Parse(input)

		if err == nil {
			ast.Name = nt.Name

			if len(result.Children) == 0 {
				ast.Children = make([]Node, 1, 1)
				ast.Children[0] = result
			} else {
				ast.Children = result.Children
			}

			return
		}
	}

	err = fmt.Errorf("Failed to parse non-terminal <%s>.", nt.Name)
	return
}

func (nt NonTerminal) String() string {
	return fmt.Sprintf("<%s>", nt.Name)
}