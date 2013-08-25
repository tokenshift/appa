package appa

type lexeme struct {
	term Terminal
	value string
}

type lexer struct {
	terms []Terminal
}

func createLexer(g *grammar) lexer {
	var lex lexer

	numTerms := len(g.literals) + len(g.regexes)
	lex.terms = make([]Terminal, 0, numTerms)

	for _, tkn := range(g.literals) {
		lex.terms = append(lex.terms, tkn)
	}

	for _, tkn := range(g.regexes) {
		lex.terms = append(lex.terms, tkn)
	}

	return lex
}

// Gets the next lexeme from the input stream.
func (lex lexer) next(in *stringBuffer) (l lexeme, ok bool) {
	if in.eof() {
		l.term = eof
		l.value = ""
		ok = true
		return
	}

	max := -1

	for _, tkn := range(lex.terms) {
		// The longest match will be used.
		match, ok := tkn.match(in)
		if ok && len(match) > max {
			l.term = tkn
			l.value = match

			max = len(match)
		}
	}

	ok = max > -1
	if ok {
		in.consume(max)
	}

	return
}
