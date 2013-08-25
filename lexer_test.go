package appa

import "strings"
import "testing"

func Test_LexingInput(t *testing.T) {
	input := createStringBuffer(strings.NewReader("123testing another func test 567.1314 hello"))

	g := NewGrammar()
	keyword := g.Lit("func")
	num := g.Regex("\\d+")
	dec := g.Regex("\\d+\\.\\d+")
	word := g.Regex("[a-z]+")
	space := g.Regex("\\s+")

	lexer := createLexer(g.(*grammar))

	lexemes := make([]lexeme, 0)
	for lexeme, ok := lexer.next(input); ok && lexeme.term != eof; lexeme, ok = lexer.next(input) {
		lexemes = append(lexemes, lexeme)
	}

	// 123
	// testing
	//
	// another
	//
	// func
	//
	// test
	//
	// 567.1314
	//
	// hello

	if !(assertIntEquals(t, 12, len(lexemes))) {
		return
	}

	assertEquals(t, num, lexemes[0].term)
	assertStringEquals(t, "123", lexemes[0].value)

	assertEquals(t, word, lexemes[1].term)
	assertStringEquals(t, "testing", lexemes[1].value)

	assertEquals(t, space, lexemes[2].term)

	assertEquals(t, word, lexemes[3].term)
	assertStringEquals(t, "another", lexemes[3].value)

	assertEquals(t, space, lexemes[4].term)

	assertEquals(t, keyword, lexemes[5].term)
	assertStringEquals(t, "func", lexemes[5].value)

	assertEquals(t, space, lexemes[6].term)

	assertEquals(t, word, lexemes[7].term)
	assertStringEquals(t, "test", lexemes[7].value)

	assertEquals(t, space, lexemes[8].term)

	assertEquals(t, dec, lexemes[9].term)
	assertStringEquals(t, "567.1314", lexemes[9].value)

	assertEquals(t, space, lexemes[10].term)

	assertEquals(t, word, lexemes[11].term)
	assertStringEquals(t, "hello", lexemes[11].value)
}
