package appa

import "regexp"

// A string literal that can act as
// either a parse rule or a node.
type Lit string

// Makes matching the rule optional.
func Opt(r Rule) Rule {
	return optional { r }
}

// Creates a rule that will match the specified
// regular expression.
func Regex(pattern string) Rule {
	return regex { regexp.MustCompile(pattern) }
}

// Concatenates the rules into a single rule.
func Seq(rules ...Rule) Rule {
	return Sequence(rules)
}

// Discards leading whitespace before attempting
// to match the rule.
func Trim(r Rule) Rule {
	if seq, ok := r.(Sequence); ok {
		// Trim applies to all children of the sequence.
		for i, rule := range(seq) {
			seq[i] = Trim(rule)
		}
		return seq
	} else {
		return trim { r }
	}
}
