package appa

import "bytes"
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
	lookahead Terminal
}

func createLALRItem(nt *nonTerminal, body []Token, pos int) lalrItem {
	return lalrItem {
		nt,
		body,
		pos,
		nil,
	}
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

// Creates a new item by incrementing the position
// of this one.
func (item lalrItem) inc() (out lalrItem, ok bool) {
	if item.pos >= len(item.body) {
		ok = false
		return
	}

	out.nt = item.nt
	out.body = item.body
	out.pos = item.pos + 1
	ok = true

	return
}

// Gets the token immediately following the parse position.
func (item lalrItem) next() Token {
	if item.pos >= len(item.body) {
		return nil
	}

	return item.body[item.pos]
}

func (item lalrItem) String() string {
	out := new(bytes.Buffer)

	fmt.Fprint(out, item.nt.String())
	fmt.Fprint(out, " →")

	for i, tkn := range (item.body) {
		if i == item.pos {
			fmt.Fprint(out, " ·")
		}

		fmt.Fprintf(out, " %v", tkn)
	}

	if item.pos == len(item.body) {
		fmt.Fprint(out, " ·")
	}

	if item.lookahead != nil {
		fmt.Fprintf(out, ", %v", item.lookahead)
	}

	return out.String()
}
