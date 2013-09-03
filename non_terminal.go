package appa

import "fmt"

type NonTerminal interface {
	Token

	Match(tokens ...Token) Rule
}

type Rule interface {
}

type nonTerminal struct {
	name string
	rules []rule

	// Tag for system-generated non-terminals.
	special string
}

func (nt *nonTerminal) Equals(other Token) bool {
	if nt2, ok := other.(*nonTerminal); ok {
		return nt2 == nt
	} else {
		return false
	}
}

func (nt *nonTerminal) first() []Terminal {
	processed := make(map[*nonTerminal]bool)
	terminals := make(map[Terminal]bool)
	queue := []*nonTerminal{ nt }


	for len(queue) > 0 {
		nonterm := queue[0]
		processed[nonterm] = true

		queue = queue[1:]

		for _, rule := range(nonterm.rules) {
			if rule.size() > 0 {
				first := rule.at(0)

				if term, ok := first.(Terminal); ok {
					terminals[term] = true
				} else {
					if _, ok := processed[first.(*nonTerminal)]; !ok {
						queue = append(queue, first.(*nonTerminal))
					}
				}
			}
		}
	}

	result := make([]Terminal, 0, len(terminals))
	for term, _ := range(terminals) {
		result = append(result, term)
	}

	return result
}

func (nt *nonTerminal) Match(tokens ...Token) Rule {
	nt.rules = append(nt.rules, createRule(tokens...))

	return nil
}

func (nt *nonTerminal) String() string {
	if nt.special == "" {
		return fmt.Sprintf("<%s>", nt.name)
	} else {
		return nt.special
	}
}
