package appa

import "regexp"

type regex struct {
	pattern *regexp.Regexp
}

func (r *regex) String() string {
	return r.pattern.String()
}
