package appa

import "fmt"
import "io"
import "strings"

// A collection of sets of LALR items connected
// by GOTOs, from which the table of parse actions
// will be constructed.
type lalrCollection struct {
	sets map[uint32][]*lalrSet
	startSet *lalrSet
}

func createLALRCollection(start *nonTerminal) (coll lalrCollection) {
	coll.sets = make(map[uint32][]*lalrSet)

	wrapper := nonTerminal {"", []rule{createRule(start)}, "{START}"}
	startItem := createLALRItem(&wrapper, wrapper.rules[0], 0)

	coll.createLALRSets(startItem)
	coll.propagateLookaheads()

	return
}

// Adds a new LALR set to this collection, if it is
// not already present.
// Returns a pointer to either the newly added set or
// the already present set.
func (coll *lalrCollection) addSet(set *lalrSet) (isNew bool, out *lalrSet) {
	hash := set.hash()
	if sets, ok := coll.sets[hash]; ok {
		for _, present := range(sets) {
			if set.equals(present) {
				return false, present
			}
		}

		coll.sets[hash] = append(sets, set)
	} else {
		coll.sets[hash] = []*lalrSet{set}
	}

	return true, set
}

// Constructs the collection of LALR sets, starting with the
// specified item.
func (coll *lalrCollection) createLALRSets(item lalrItem) {
	item.addLookahead(eof)
	startSet := createLALRSet(item)

	coll.addSet(startSet)
	coll.startSet = startSet

	// Keep a queue of sets that have not yet been processed.
	newSets := []*lalrSet{coll.startSet}

	for len(newSets) > 0 {
		set := newSets[0]
		newSets = newSets[1:]

		// Compute the closure of the set.
		closure := set.closure()

		// Determine what 'next' tokens are present.
		// For each 'next' token:
		for _, tkn := range(closure.nextTokens()) {
			// Create the kernel of the GOTO set.
			gto := closure.createGoto(tkn)

			// Add the kernel to this collection.
			isNew, gto := coll.addSet(gto)

			// Add GOTO to this set.
			set.addGoto(tkn, gto)

			// Add the new set to the queue.
			if isNew {
				newSets = append(newSets, gto)
			}
		}
	}
}

// Creates a table of parse actions from this LALR collection.
func (coll *lalrCollection) createTable() (table actionTable) {
	return
}

// Iterates through each of the LALR sets in this collection.
func (coll *lalrCollection) each(f func(*lalrSet)) {
	for _, sets := range(coll.sets) {
		for _, set := range(sets) {
			f(set)
		}
	}
}

// Computes lookaheads for the LALR set kernels in this collection.
func (coll *lalrCollection) propagateLookaheads() {
	// Determine spontaneously generated lookaheads
	// and propagation paths.

	// For each kernel K in the collection
	coll.each(func (set *lalrSet) {
		// For each item A → α·β in K
		set.each(func (item lalrItem) {
			// Let J be CLOSURE({[A → α·β, #]})
			i2 := createLALRItem(item.nt, item.body, item.pos, bogy)

			j := createLALRSet(i2).closure()

			j.each(func (cItem lalrItem) {
				if cItem.hasLookahead(bogy) {
					// If [B → γ·Xδ, #] is in J, conclude that
					// lookaheads propagate from A → α·β in K
					// to B → γX·δ in GOTO(K, X).
				} else {
					// If [B → γ·Xδ, a] is in J and 'a' is not #,
					// conclude that lookahead 'a' is generated
					// spontaneously for item B → γX·δ in GOTO(K, X).
				}
			})
		})
	})
}

// Gets the number of LALR sets in the collection.
func (coll *lalrCollection) size() (count int) {
	for _, list := range(coll.sets) {
		count = count + len(list)
	}
	return
}

// Writes the collection of LALR sets in DOT format
// to the specified output stream.
func (coll *lalrCollection) writeTo(out io.Writer) {
	setNumbers := make(map[*lalrSet]int)

	fmt.Fprint(out, "digraph LALR_Collection {\n\tnode[shape=box];\n")

	for _, sets := range(coll.sets) {
		for _, set := range(sets) {
			num, ok := setNumbers[set]
			if !ok {
				num = len(setNumbers) + 1
				setNumbers[set] = num
			}

			fmt.Fprintf(out, "\n\t%d [", num)
			if set == coll.startSet {
				fmt.Fprint(out, "penwidth=2\n\t")
			}
			fmt.Fprint(out, "label=<<b>")

			// Write kernel items.
			for _, items := range(set.items) {
				for _, item := range(items) {
					itemString := strings.Replace(item.String(), "<", "&amp;lt;", -1)
					itemString = strings.Replace(itemString, ">", "&amp;gt;", -1)
					fmt.Fprintf(out, "\n\t%s<br align=\"left\" />", itemString)
				}
			}

			// Write closure (non-kernel) items.
			closure := set.closure()
			if closure.size() > set.size() {
				fmt.Fprint(out, "\n\t</b><i>")
				for _, items := range(closure.items) {
					for _, item := range(items) {
						// Only write if wasn't in kernel.
						if !set.contains(item) {
							itemString := strings.Replace(item.String(), "<", "&amp;lt;", -1)
							itemString = strings.Replace(itemString, ">", "&amp;gt;", -1)
							fmt.Fprintf(out, "\n\t%s<br align=\"left\" />", itemString)
						}
					}
				}

				fmt.Fprint(out, "\n\t</i>>];\n")
			} else {
				fmt.Fprint(out, "\n\t</b>>];\n")
			}

			// Write gotos for this set.
			//gotos map[Token]*lalrSet
			for tkn, set2 := range(set.gotos) {
				num2, ok := setNumbers[set2]
				if !ok {
					num2 = len(setNumbers) + 1
					setNumbers[set2] = num2
				}

				label := strings.Replace(fmt.Sprint(tkn), "\"", "\\\"", -1)
				fmt.Fprintf(out, "\n\t%d -> %d [label=\"%s\"];", num, num2, label)
			}

			if len(set.gotos) > 0 {
				fmt.Fprint(out, "\n")
			}
		}
	}

	fmt.Fprint(out, "}\n")
}
