package appa

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
