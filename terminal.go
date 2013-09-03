package appa

import "fmt"
import "regexp"

type Terminal interface {
	Token

	match(in *stringBuffer) (match string, ok bool)
}

// Symbol representing a terminal that will
// never otherwise occur in a grammar.
type _bogy int
var bogy _bogy = -1

func (_bogy) Equals(other Token) bool {
	_, ok := other.(_bogy)
	return ok
}

func (b _bogy) first() []Terminal {
	return []Terminal{b}
}

func (_bogy) match(in *stringBuffer) (match string, ok bool) {
	return "", false
}

func (_bogy) String() string {
	return "#"
}

// Represents the end-of-file marker.
type _eof int
var eof _eof = 0

func (_eof) Equals(other Token) bool {
	_, ok := other.(_eof)
	return ok
}

func (e _eof) first() []Terminal {
	return []Terminal{e}
}

func (_eof) match(in *stringBuffer) (match string, ok bool) {
	match = ""
	ok = in.eof()
	return
}

func (_eof) String() string {
	return "$"
}

// Matches a string literal.
type lit string

func (l lit) Equals(other Token) bool {
	if l2, ok := other.(lit); ok {
		return string(l2) == string(l)
	} else {
		return false
	}
}

func (l lit) first() []Terminal {
	return []Terminal{l}
}

func (l lit) match(in *stringBuffer) (match string, ok bool) {
	ok = in.readLiteral(string(l), 0)
	if ok {
		match = string(l)
	}

	return
}

func (l lit) String() string {
	return fmt.Sprintf("\"%s\"", string(l))
}

// Matches a regular expression.
type regex struct {
	pattern *regexp.Regexp
}

func (r regex) Equals(other Token) bool {
	if r2, ok := other.(regex); ok {
		return r2.pattern == r.pattern
	} else {
		return false
	}
}

func (r regex) first() []Terminal {
	return []Terminal{r}
}

func (r regex) match(in *stringBuffer) (match string, ok bool) {
	ok, match = in.readPattern(r.pattern, 0)
	return
}

func (r regex) String() string {
	return fmt.Sprintf("/%v/", r.pattern)
}
