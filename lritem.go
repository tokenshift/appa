package appa

import "bytes"

// An LR(*) item, combining a specific
// production and a parse position.
type lrItem struct {
	prod *production
	pos int
}

func (item lrItem) next() Token {
	if item.pos >= item.prod.length() {
		return nil;
	}

	return item.prod.body[item.pos]
}

func (item lrItem) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(item.prod.nt.String())
	buffer.WriteString(" →")

	for i, tkn := range(item.prod.body) {
		if i == item.pos {
			buffer.WriteString(" ·")
		}

		buffer.WriteString(" ")
		buffer.WriteString(tkn.String())
	}

	if item.pos == len(item.prod.body) {
		buffer.WriteString(" ·")
	}

	return buffer.String()
}

// A collection of LR items.
type itemSet struct {
	items map[lrItem]bool
}

func (set itemSet) add(item lrItem) {
	set.items[item] = true
}


func (set itemSet) closure() itemSet {
	return set
}

func (set itemSet) remove(item lrItem) {
	delete(set.items, item)
}
