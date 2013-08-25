package appa

import "bufio"
import "fmt"
import "io"
import "regexp"
import "strings"

type stringBuffer struct {
	buffer string
	isEof bool
	input *bufio.Reader
}

// Creates a new string buffer from the specified input stream.
func createStringBuffer(input io.Reader) *stringBuffer {
	buffer := new(stringBuffer)
	buffer.isEof = false
	buffer.input = bufio.NewReader(input)
	return buffer
}

// Consumes N characters from the input stream.
func (b *stringBuffer) consume(n int) string {
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

// Discards N characters from the input stream.
// Returns the number of characters actually discarded.
func (b *stringBuffer) discard(n int) int {
	b.prepare(n)

	if n < len(b.buffer) {
		b.buffer = b.buffer[n:]
	} else {
		n = len(b.buffer)
		b.buffer = ""
	}

	return n
}

// Whether the end of the input has been reached.
func (b *stringBuffer) eof() bool {
	return b.isEof && len(b.buffer) == 0
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

// Attempts to match a string literal.
func (b *stringBuffer) readLiteral(text string, offset int) (bool) {
	b.prepare(offset + len(text))

	if len(b.buffer) < len(text) + offset {
		return false
	}

	if strings.HasPrefix(b.buffer[offset:], text) {
		return true
	}

	return false
}

// Attempts to match a regular expression.
func (b *stringBuffer) readPattern(pattern *regexp.Regexp, offset int) (bool, string) {
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
