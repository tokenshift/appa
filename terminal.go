package appa

import "fmt"
import "regexp"
import "strings"

type Terminal interface {
	Token
	Match(string) string
}

type lit struct { string }

func (l lit) firsts() []Terminal {
	return []Terminal { l }
}

func (l lit) Match(input string) string {
	if strings.HasPrefix(input, l.string) {
		return l.string
	} else {
		return ""
	}
}

func (l lit) String() string {
	return l.string
}

type reg struct { *regexp.Regexp }

func (r reg) firsts() []Terminal {
	return []Terminal { r }
}

func (r reg) Match(input string) string {
	match := r.FindStringSubmatchIndex(input)
	if match != nil && match[0] == 0 {
		return input[match[0]:match[1]]
	} else {
		return ""
	}
}

func (r reg) String() string {
	return fmt.Sprintf("/%v/", r.Regexp)
}

type eof struct {}

func (eof) firsts() []Terminal {
	return []Terminal { eof {} }
}

func (eof) Match(input string) string {
	return ""
}

func (eof) String() string {
	return "$"
}
