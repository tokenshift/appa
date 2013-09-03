package appa

// A terminal or non-terminal; anything
// that can appear in the body of a production.
type Token interface {
	Equals(other Token) bool

	// Computes the FIRST set for the token.
	first() []Terminal
}
