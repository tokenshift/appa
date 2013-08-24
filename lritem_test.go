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

func Test_LRItemGetNextToken(t *testing.T) {
	nt := new(nonTerminal)
	nt.name = "FOO"

	prod := new(production)
	prod.nt = nt
	prod.body = []Token{
		&lit{"Test"},
		&lit{"Foo"},
		&lit{"Bar"},
	}

	var item lrItem
	var tkn Token
	var l *lit
	var ok bool

	item = lrItem { prod, 0 }
	tkn = item.next()
	if l, ok = tkn.(*lit); !ok {
		t.Errorf("Expected a literal.")
	}
	assertStringEquals(t, "Test", l.String())

	item = lrItem { prod, 1 }
	tkn = item.next()
	if l, ok = tkn.(*lit); !ok {
		t.Errorf("Expected a literal.")
	}
	assertStringEquals(t, "Foo", l.String())

	item = lrItem { prod, 2 }
	tkn = item.next()
	if l, ok = tkn.(*lit); !ok {
		t.Errorf("Expected a literal.")
	}
	assertStringEquals(t, "Bar", l.String())

	item = lrItem { prod, 3 }
	tkn = item.next()
	assertNil(t, tkn)
}
