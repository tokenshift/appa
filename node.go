package appa

import "bytes"
import "fmt"

var Empty NodeList = make([]Node, 0, 0)

// Represents a node in an abstract
// synax tree.
type Node interface {
	// The text value of the node, or nil.
	Val() fmt.Stringer

	// Any children of the node.
	Children() []Node
}

// A sequence of parsed nodes.
type NodeList []Node

func (nodes NodeList) Children() []Node {
	return nodes
}

func (nodes NodeList) Val() fmt.Stringer {
	return nil
}

// A node with a name and children.
// Used to represent non-terminals.
type Named struct {
	Name string
	children []Node
}

// Creates a new node with the specified
// name and children.
func NamedNode(name string, children ...Node) Named {
	return Named { name, children }
}

func (n Named) Children() []Node {
	return n.children
}

func (n Named) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("(")
	buffer.WriteString(n.Name)

	for _, node := range(n.children) {
		buffer.WriteString(" ")
		if node == nil {
			buffer.WriteString("()")
		} else {
			buffer.WriteString(fmt.Sprint(node))
		}
	}

	buffer.WriteString(")")

	return buffer.String()
}

func (n Named) Val() fmt.Stringer {
	return Lit(n.Name)
}
