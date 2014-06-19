package appa

import "bytes"

type item struct {
	head NonTerminal
	body []Token
	pos int
	lookahead Terminal
}

func (it item) String() string {
	var buf bytes.Buffer

	buf.WriteString(it.head.String())
	buf.WriteString(" →")

	for i, tkn := range(it.body) {
		if i == it.pos {
			buf.WriteString(" ·")
		}
		buf.WriteString(" ")
		buf.WriteString(tkn.String())
	}

	buf.WriteString(", ")
	buf.WriteString(it.lookahead.String())

	return buf.String()
}

func (it item) eq(it2 item) bool {
	if it.head != it2.head {
		return false
	}

	if it.pos != it2.pos {
		return false
	}

	if len(it.body) != len(it2.body) {
		return false
	}

	for i, t := range(it.body) {
		if t != it2.body[i] {
			return false
		}
	}

	return true
}

func (it item) next() Token {
	if it.pos < len(it.body) {
		return it.body[it.pos]
	} else {
		return nil
	}
}
