package appa

import "fmt"
import "regexp"

type Terminal interface {
	Token

	match(in *stringBuffer) (match string, ok bool)
}

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
