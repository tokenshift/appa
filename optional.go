package appa

type optional struct {
	rule Rule
}

func (o optional) And(r Rule) Rule {
	return makeSequence(o, r)
}

func (o optional) Match(buffer StringBuffer, offset int) int {
	matched := o.rule.Match(buffer, offset)
	
	if matched > 0 {
		return matched
	}

	return 0
}

func (o optional) Parse(buffer StringBuffer) (node Node, err error) {
	matched := o.rule.Match(buffer, 0)
	if matched > 0 {
		node, err = o.rule.Parse(buffer)
	} else {
		node = Node {
			"",
			make([]Node, 0, 0),
		}
	}

	return
}



func (r Literal) Optional() Rule {
	return optional { r }
}

func (r NonTerminal) Optional() Rule {
	return optional { r }
}

func (r optional) Optional() Rule {
	return r
}

func (r Regex) Optional() Rule {
	return optional { r }
}

func (r sequence) Optional() Rule {
	return optional { r }
}
