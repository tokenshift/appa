package appa

import "fmt"

type eof struct {}

func (eof) Match(buffer StringBuffer, offset int) int {
	if buffer.EofAt(offset) {
		return 0
	} else {
		return -1
	}
}

func (eof) Parse(buffer StringBuffer) ([]Node, error) {
	if buffer.Eof() {
		return []Node{}, nil
	} else {
		return nil, fmt.Errorf("Did not find EOF.")
	}
}

func (eof) String() string {
	return "EOF"
}
