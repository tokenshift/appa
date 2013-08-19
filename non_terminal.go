package appa

import (
	"fmt"
)

type NonTerminal struct {
	Name string
	rules []Rule
}

func (g *grammar) NonTerminal(name string) (nt NonTerminal) {
	nt, ok := g.nonterminals[name]
	if !ok {
		nt = NonTerminal {
			name,
			make([]Rule, 0),
		}
		g.nonterminals[name] = nt
	}
	return
}

func (nt *NonTerminal) AddRule(r Rule) {
	nt.rules = append(nt.rules, r)
}

func (nt NonTerminal) And(r Rule) Rule {
	rules := make([]Rule, 2, 2)
	rules[0] = nt
	rules[1] = r
	return sequence {
		rules,
	}
}

func (nt NonTerminal) ParseString(input string) (ast Node, err error) {
	for _, rule := range nt.rules {
		var result Node
		result, err = rule.ParseString(input)
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

	err = fmt.Errorf("Failed to parse non-terminal <{0}>.", nt.Name)
	return
}
