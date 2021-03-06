package appa

import "bytes"
import "fmt"

// A set of LALR items.
type lalrSet struct {
	items map[uint32][]lalrItem
	gotos map[Token]*lalrSet
}

// Create a new LALR set containing the specified items.
func createLALRSet(items ...lalrItem) (set *lalrSet) {
	set = new(lalrSet)
	set.items = make(map[uint32][]lalrItem)
	set.gotos = make(map[Token]*lalrSet)

	for _, item := range(items) {
		set.addItem(item)
	}

	return
}

// Adds a GOTO (transition on a specific token) to this set.
func (set *lalrSet) addGoto(tkn Token, other *lalrSet) {
	set.gotos[tkn] = other
}

// Adds a new item to the LALR set.
// Returns true if the item was not already present.
func (set *lalrSet) addItem(item lalrItem) bool {
	hash := item.hash()
	if items, ok := set.items[hash]; ok {
		for _, present := range(items) {
			if item.equals(present) {
				// Item is already present.
				return false
			}
		}

		set.items[hash] = append(items, item)
	} else {
		set.items[hash] = []lalrItem{item}
	}

	return true
}

// Computes the closure of this LALR set.
func (set *lalrSet) closure() (out *lalrSet) {
	out = createLALRSet()

	// Keep a queue of items that still need
	// to be processed.
	newItems := make([]lalrItem, 0)

	for _, items := range(set.items) {
		for _, item := range(items) {
			if out.addItem(item) {
				newItems = append(newItems, item)
			}
		}
	}

	for len(newItems) > 0 {
		item := newItems[0]
		newItems = newItems[1:]

		if next := item.next(); next != nil {
			if nt, ok := next.(*nonTerminal); ok {
				for _, rule := range(nt.rules) {
					var lookaheads []Terminal
					if item.pos < item.body.size() {
						lookaheads = item.body.at(item.pos).first()
					} else {
						lookaheads = item.lookaheads
					}

					it2 := createLALRItem(nt, rule, 0, lookaheads...)
					if out.addItem(it2) {
						newItems = append(newItems, it2)
					}
				}
			}
		}
	}

	return
}

// Creates the kernel of an LALR set resulting from a
// transition on the specified token.
func (set *lalrSet) createGoto(on Token) (out *lalrSet) {
	out = createLALRSet()

	for _, items := range(set.items) {
		for _, item := range(items) {
			if item.next() == on {
				if next, ok := item.inc(); ok {
					out.addItem(next)
				}
			}
		}
	}

	return
}

// Checks whether this set contains the specified LALR item.
func (set *lalrSet) contains(item lalrItem) bool {
	hash := item.hash()

	if items, ok := set.items[hash]; ok {
		for _, item2 := range(items) {
			if item.equals(item2) {
				return true
			}
		}
	}

	return false
}

// Iterates through the items in this LALR set.
func (set *lalrSet) each(f func(lalrItem)) {
	for _, items := range(set.items) {
		for _, item := range(items) {
			f(item)
		}
	}
}

// Value equality for LALR sets.
func (set *lalrSet) equals(other *lalrSet) bool {
	if set == other {
		return true
	}

	if set.size() != other.size() {
		return false
	}

	for _, list := range(set.items) {
		for _, item := range(list) {
			if !other.contains(item) {
				return false
			}
		}
	}

	return true
}

// Hashing function for LALR set lookup.
func (set *lalrSet) hash() (val uint32) {
	// Need hash to be order independent, so
	// just XOR everything.
	for _, list := range(set.items) {
		for _, item := range(list) {
			val = val ^ item.hash()
		}
	}

	return
}

// Gets a list of 'next' tokens for all of the
// items in this set.
func (set *lalrSet) nextTokens() []Token {
	tokens := make([]Token, 0)

	for _, items := range(set.items) {
		for _, item := range(items) {
			if next := item.next(); next != nil {
				exists := false
				for _, tkn := range(tokens) {
					if next.Equals(tkn) {
						exists = true
						break
					}
				}
				if !exists {
					tokens = append(tokens, next)
				}
			}
		}
	}

	return tokens
}

// The number of items in the set.
func (set *lalrSet) size() (count int) {
	for _, list := range(set.items) {
		count = count + len(list)
	}
	return
}

func (set *lalrSet) String() string {
	out := new(bytes.Buffer)

	first := true
	for _, items := range(set.items) {
		for _, item := range(items) {
			if !first {
				fmt.Fprint(out, "\n")
			}
			fmt.Fprint(out, item)
			first = false
		}
	}

	closure := set.closure()
	for _, items := range(closure.items) {
		for _, item := range(items) {
			if !set.contains(item) {
				fmt.Fprint(out, "\n- %v", item)
			}
		}
	}

	return out.String()
}
