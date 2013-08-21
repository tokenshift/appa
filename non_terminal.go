package appa

import "fmt"

type nonTerminal struct {
	g *grammar
	name string
}

func (nt nonTerminal) AddRule(r Rule) {
	nt.g.rules[nt.name] = append(nt.g.rules[nt.name], r)
}

func (nt nonTerminal) Match(input StringBuffer, offset int) int {
	for _, rule := range nt.g.rules[nt.name] {
		if matched := rule.Match(input, offset); matched > 0 {
			return matched
		}
	}

	return -1
}

func (nt nonTerminal) Name() string {
	return nt.name
}

func (nt nonTerminal) Parse(input StringBuffer) (ast Node, err error) {
	for _, rule := range nt.g.rules[nt.name] {
		var result Node
		result, err = rule.Parse(input)

		if err == nil {
			// TODO: generalize this branch into a reduction rule
			// associated with the node.
			if seq, ok := result.(NodeList); ok {
				// Flatten the sequence into the non-terminal.
				ast = NamedNode(nt.name, seq.Children()...)
			} else {
				// Use the node as is.
				ast = NamedNode(nt.name, result)
			}
			return
		}
	}

	err = fmt.Errorf("Failed to parse non-terminal <%s>.", nt.Name)
	return
}

func (nt nonTerminal) String() string {
	return fmt.Sprintf("<%s>", nt.Name)
}
