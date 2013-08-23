package appa

import "fmt"

// A terminal or non-terminal; anything
// that can appear in the body of a production.
type Token interface {
	fmt.Stringer
}
