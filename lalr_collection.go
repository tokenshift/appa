package appa

// A collection of sets of LALR items connected
// by GOTOs, from which the table of parse actions
// will be constructed.
type lalrCollection struct {
	sets map[int][]lalrSet
}

func createLALRCollection(g *grammar) (coll lalrCollection) {
	coll.sets = make(map[int][]lalrSet, len(g.nonterminals))
	return
}

func (coll lalrCollection) createTable() (table actionTable) {
	return
}
