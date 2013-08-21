package appa

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Provides buffered text reading with lookahead.
type StringBuffer interface {
	// Consumes N characters from the input stream.
	Consume(n int) string

	// Discards N characters from the input stream.
	// Returns the number of characters actually discarded.
	Discard(n int) int

	// Whether the end of the input has been reached.
	Eof() bool

	// Whether the end of the input is at the specified offset.
	EofAt(offset int) bool

	// Attempts to match a string literal.
	ReadLiteral(text string, offset int) (ok bool)

	// Attempts to match a regular expression.
	ReadPattern(pattern *regexp.Regexp, offset int) (ok bool, match string)
}

func CreateStringBuffer(input io.Reader) StringBuffer {
	buffer := new(stringBuffer)
	buffer.isEof = false
	buffer.input = bufio.NewReader(input)
	return buffer
}

type stringBuffer struct {
	buffer string
	isEof bool
	input *bufio.Reader
}

func (b *stringBuffer) Consume(n int) string {
	b.prepare(n)

	var str string
	if n < len(b.buffer) {
		str = b.buffer[:n]
		b.buffer = b.buffer[n:]
	} else {
		str = b.buffer
		b.buffer = ""
	}

	return str
}

func (b *stringBuffer) Discard(n int) int {
	b.prepare(n)

	if n < len(b.buffer) {
		b.buffer = b.buffer[n:]
	} else {
		n = len(b.buffer)
		b.buffer = ""
	}

	return n
}

func (b *stringBuffer) Eof() bool {
	return b.isEof && len(b.buffer) == 0
}

func (b *stringBuffer) EofAt(offset int) bool {
	b.prepare(offset + 1)
	return b.isEof && offset >= len(b.buffer)
}

func (b *stringBuffer) ReadLiteral(text string, offset int) (bool) {
	b.prepare(offset + len(text))

	if len(b.buffer) < len(text) + offset {
		return false
	}

	if strings.HasPrefix(b.buffer[offset:], text) {
		return true
	}

	return false
}

func (b *stringBuffer) ReadPattern(pattern *regexp.Regexp, offset int) (bool, string) {
	b.prepare(256 + offset)

	match := pattern.FindStringIndex(b.buffer[offset:])
	if match == nil || match[0] != 0 {
		return false, ""
	}

	return true, b.buffer[match[0] + offset:match[1] + offset]
}

func (b *stringBuffer) String() string {
	return fmt.Sprintf("\"%s\"", b.buffer)
}

// Loads up to N characters into the input buffer.
func (b *stringBuffer) prepare(n int) {
	for !b.isEof && len(b.buffer) < n {
		text, err := b.input.ReadString('\n')
		b.buffer = b.buffer + text
		if err == io.EOF {
			b.isEof = true
		}
	}
}
