package appa

import (
	"bytes"
)

type Node struct {
	Name string
	Children []Node
}

func (n Node) String() string {
	var buffer bytes.Buffer

	if len(n.Children) > 0 {
		buffer.WriteString("(")
	}

	buffer.WriteString(n.Name)

	for _, child := range(n.Children) {
		buffer.WriteString(" ")
		buffer.WriteString(child.String())
	}

	if len(n.Children) > 0 {
		buffer.WriteString(")")
	}

	return buffer.String()
}
