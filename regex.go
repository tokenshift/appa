package appa

import "fmt"
import "regexp"

type regex struct {
	pattern *regexp.Regexp
}

func (r *regex) String() string {
	return fmt.Sprintf("/%v/", r.pattern)
}
