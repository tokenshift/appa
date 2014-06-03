package appa

import "fmt"
import "regexp"
import "strings"

type Terminal interface {
	Token
	Match(string) string
}

type lit struct { string }

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

func (r reg) Match(input string) string {
	match := r.FindStringSubmatchIndex(input)
	if match != nil && match[0] == 0 {
		return input[match[0]:match[1]]
	} else {
		return ""
	}
}

func (r reg) String() string {
	return fmt.Sprintf("/%v/", r)
}
