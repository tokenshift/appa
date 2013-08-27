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
	lookahead Terminal
}

func createLALRItem(nt *nonTerminal, body rule, pos int) lalrItem {
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


	if !item.body.equals(other.body) {
		return false
	}

	if item.lookahead != other.lookahead {
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

	if item.lookahead != nil {
		io.WriteString(hash, fmt.Sprint(item.lookahead))
	}

	return hash.Sum32()
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

	if item.lookahead != nil {
		fmt.Fprintf(out, ", %v", item.lookahead)
	}

	return out.String()
}
