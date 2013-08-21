package appa

// A parsing rule that will be used to
// match input text.
type Rule interface {
	// Checks whether this rule matches the start of the input stream.
	// Returns the number of characters matched, or -1 on failure.
	Match(input StringBuffer, offset int) int

	// Parses and consumes the matched portion of the input.
	Parse(input StringBuffer) (Node, error)
}
