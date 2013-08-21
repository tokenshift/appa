package appa

import "fmt"

type nonTerminal struct {
	g *grammar
	name string
}

// Combines a parse rule with a set of reductions.
type rulePackage struct {
	rule Rule
	reduction Reduce
}

func (nt nonTerminal) AddReduction(r Rule, reduction Reduce) {
	if reduction == nil {
		reduction = defaultReduction(nt)
	}

	rp := rulePackage {
		r,
		reduction,
	}

	nt.g.rules[nt.name] = append(nt.g.rules[nt.name], rp)
}

func (nt nonTerminal) AddRule(r Rule) {
	nt.AddReduction(r, nil)
}

func (nt nonTerminal) Match(input StringBuffer, offset int) int {
	for _, rp := range nt.g.rules[nt.name] {
		rule := rp.rule
		if matched := rule.Match(input, offset); matched > 0 {
			return matched
		}
	}

	return -1
}

func (nt nonTerminal) Name() string {
	return nt.name
}

func (nt nonTerminal) Parse(input StringBuffer) (result []Node, err error) {
	for _, rp := range nt.g.rules[nt.name] {
		rule := rp.rule

		result, err = rule.Parse(input)

		if err == nil {
			result = []Node{reduce(nt, rp, result)}
			return
		}
	}

	err = fmt.Errorf("Failed to parse non-terminal <%s>.", nt.Name)
	return
}

func (nt nonTerminal) String() string {
	return fmt.Sprintf("<%s>", nt.name)
}

// Returns the matched nodes as children of a named node.
func defaultReduction(nt nonTerminal) Reduce {
	return func(matched []Node) Node {
		return NodeNamed(nt.name, matched...)
	}
}

// Reduces the parsed content using the associated reduction rule.
func reduce(nt nonTerminal, rp rulePackage, matched []Node) Node {
	return rp.reduction(matched)
}
