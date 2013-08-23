package appa

import "testing"

func Test_LRItemToString(t *testing.T) {
	nt := new(nonTerminal)
	nt.name = "FOO"

	prod := new(production)
	prod.nt = nt
	prod.body = []Token{
		&lit{"Test"},
		&lit{"Foo"},
		&lit{"Bar"},
	}

	item := lrItem { prod, 0 }
	assertStringEquals(t, "<FOO> → · Test Foo Bar", item.String())

	item = lrItem { prod, 1 }
	assertStringEquals(t, "<FOO> → Test · Foo Bar", item.String())

	item = lrItem { prod, 2 }
	assertStringEquals(t, "<FOO> → Test Foo · Bar", item.String())

	item = lrItem { prod, 3 }
	assertStringEquals(t, "<FOO> → Test Foo Bar ·", item.String())
}
