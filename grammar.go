package appa

type Grammar interface {
	Literal(value string) (lit Literal)
	NonTerminal(name string) (nt NonTerminal)
	Regex(pattern string) (rx Regex, err error)
}

func NewGrammar() Grammar {
	g := new(grammar)

	g.literals = make(map[string]Literal)
	g.nonterminals = make(map[string]NonTerminal)
	g.regexes = make(map[string]Regex)

	g.rules = make(map[string][]Rule)

	return g
}

// A parsing rule that will be used to
// match input text.
type Rule interface {
	// Constructs a rule consisting of the concatenation
	// of this rule and the other.
	And(r Rule) Rule

	// Checks whether this rule matches the start of the input stream.
	// Returns the number of characters matched, or -1 on failure.
	Match(input StringBuffer, offset int) int

	// Makes the rule optional.
	Optional() Rule

	// Parses and consumes the matched portion of the input.
	Parse(input StringBuffer) (Node, error)

	// Ignores/discards leading whitespace when attempting to match
	// this rule.
	Trim() Rule
}

type grammar struct {
	literals map[string]Literal
	nonterminals map[string]NonTerminal
	regexes map[string]Regex

	rules map[string][]Rule
}
