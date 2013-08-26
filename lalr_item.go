package appa

import "fmt"
import "hash/fnv"
import "io"

// A single LALR item, with lookaheads.
type lalrItem struct {
	// The head of the production.
	nt *nonTerminal
	// The body of the production.
	body []Token
	// The position of the parse in the production.
	pos int

	// Lookaheads for the item.
	lookaheads []Terminal
}

// Value equality for the LALR item.
func (item lalrItem) equals(other lalrItem) bool {
	if item.hash() != other.hash() {
		return false
	}

	if item.nt != other.nt {
		return false
	}

	if item.pos != other.pos {
		return false
	}

	if len(item.body) != len(other.body) {
		return false
	}

	for i, tkn := range(item.body) {
		if tkn != other.body[i] {
			return false;
		}
	}

	// Lookaheads are ignored in equality comparison.

	return true
}

// Hash function for LALR item lookup.
func (item lalrItem) hash() uint32 {
	hash := fnv.New32()

	io.WriteString(hash, item.nt.String())
	for _, tkn := range(item.body) {
		io.WriteString(hash, fmt.Sprint(tkn))
	}

	io.WriteString(hash, fmt.Sprint(item.pos))

	return hash.Sum32()
}
