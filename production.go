package appa

import "bytes"
import "fmt"

// A single production rule for the grammar.
type production struct {
	nt *nonTerminal
	body []Token
}

func (p *production) length() int {
	return len(p.body)
}

func (p *production) at(pos int) lrItem {
	if pos > p.length() {
		panic(fmt.Sprintf("Internal error: cannot create lritem at pos %d in production %v.", pos, p))
	}

	return lrItem {
		p,
		pos,
	}
}

func (p *production) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(p.nt.String())
	buffer.WriteString(" →")

	for _, tkn := range(p.body) {
		buffer.WriteString(" ")
		buffer.WriteString(tkn.String())
	}

	return buffer.String()
}
