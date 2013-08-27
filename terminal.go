package appa

import "fmt"
import "regexp"

type Terminal interface {
	Token

	match(in *stringBuffer) (match string, ok bool)
}

type _eof int
var eof _eof = 0

func (_eof) match(in *stringBuffer) (match string, ok bool) {
	match = ""
	ok = in.eof()
	return
}

func (_eof) String() string {
	return "$"
}

type lit string

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

func (r regex) match(in *stringBuffer) (match string, ok bool) {
	ok, match = in.readPattern(r.pattern, 0)
	return
}

func (r regex) String() string {
	return fmt.Sprintf("/%v/", r.pattern)
}
