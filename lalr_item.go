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
	body rule
	// The position of the parse in the production.
	pos int

	// Lookaheads for the item.
	lookaheads []Terminal
}

func createLALRItem(nt *nonTerminal, body rule, pos int, lookaheads ...Terminal) lalrItem {
	return lalrItem {
		nt,
		body,
		pos,
		lookaheads,
	}
}

// Adds a lookahead to this LALR item.
// Returns true if the lookahead was not already present.
func (item lalrItem) addLookahead(la Terminal) bool {
	if item.hasLookahead(la) {
		return false
	} else {
		item.lookaheads = append(item.lookaheads, la)
		return true
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


	if !item.body.equals(other.body) {
		return false
	}

	return true
}

// Hash function for LALR item lookup.
func (item lalrItem) hash() uint32 {
	hash := fnv.New32()

	io.WriteString(hash, item.nt.String())
	io.WriteString(hash, fmt.Sprint(item.body.hash()))
	io.WriteString(hash, fmt.Sprint(item.pos))

	return hash.Sum32()
}

// Checks whether the specified lookahead is present.
func (item lalrItem) hasLookahead(la Terminal) bool {
	for _, tkn := range(item.lookaheads) {
		if tkn == la {
			return true
		}
	}

	return false
}

// Creates a new item by incrementing the position
// of this one.
func (item lalrItem) inc() (out lalrItem, ok bool) {
	if item.pos >= item.body.size() {
		ok = false
		return
	}

	out.nt = item.nt
	out.body = item.body
	out.pos = item.pos + 1
	out.lookaheads = make([]Terminal, len(item.lookaheads))
	copy(out.lookaheads, item.lookaheads)
	ok = true

	return
}

// Gets the token immediately following the parse position.
func (item lalrItem) next() Token {
	if item.pos >= item.body.size() {
		return nil
	}

	return item.body.at(item.pos)
}

func (item lalrItem) String() string {
	out := new(bytes.Buffer)

	fmt.Fprint(out, item.nt.String())
	fmt.Fprint(out, " →")

	for i, tkn := range (item.body.body) {
		if i == item.pos {
			fmt.Fprint(out, " ·")
		}

		fmt.Fprintf(out, " %v", tkn)
	}

	if item.pos == item.body.size() {
		fmt.Fprint(out, " ·")
	}


	for i, la := range(item.lookaheads) {
		if i == 0 {
			fmt.Fprint(out, ",")
		}

		fmt.Fprintf(out, " %v", la)
	}

	return out.String()
}
