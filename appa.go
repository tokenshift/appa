package appa

import "fmt"

// Represents a self-contained Appa grammar,
// with a set of non-terminals and parse rules.
type Grammar interface {
	// Creates or retrieves a non-terminal with the specified name.
	NonTerminal(name string) NonTerminal

	// Creates a rule that will match the exact text specified.
	Lit(value string) Rule

	// Creates a rule that will match the specified regular expression.
	Regex(pattern string) Rule
}

// Creates and initializes a new grammar.
func NewGrammar() Grammar {
	g := new(grammar)

	g.nonterminals = make(map[string]NonTerminal)
	g.rules = make(map[string][]rulePackage)

	return g
}

// A node in a parsed abstract syntax tree.
type Node interface {
	fmt.Stringer

	// The text value of the node, or nil.
	Val() fmt.Stringer

	// Any children of the node.
	Children() []Node
}

// A non-terminal in the grammar, with associated parse rules.
type NonTerminal interface {
	// Adds a parsing rule to the non-terminal.
	// Each of the rules is tested in the order they were added;
	// the first successful match will be used.
	//
	// r: The rule to add.
	// reduction: A transformation to be applied when the rule is matched.
	AddReduction(r Rule, reductions Reduce)

	// Adds a parsing rule to the non-terminal.
	// Each of the rules is tested in the order they were added;
	// the first successful match will be used.
	AddRule(r Rule)

	// The name of the non-terminal.
	// This will be unique in the grammar.
	Name() string

	// A non-terminal itself counts as a rule, and can be associated
	// with other non-terminals (or recursively within itself).
	Rule
}

// A reduction rule that will be applied to the result
// of a successful parse rule evaluation.
//
// matched: The parsed content.
//
// result: The transformed value of this rule.
// ok: Whether the transformation was successful.
type Reduce func(matched []Node) (result Node)

// A parsing rule that will be used to
// match input text.
type Rule interface {
	fmt.Stringer

	// Checks whether this rule matches the start of the input stream.
	// Returns the number of characters matched, or -1 on failure.
	Match(input StringBuffer, offset int) int

	// Parses and consumes the matched portion of the input.
	Parse(input StringBuffer) ([]Node, error)
}
