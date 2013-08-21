package appa

import "bytes"
import "fmt"

var Empty NodeList = make([]Node, 0, 0)

// A sequence of parsed nodes.
type NodeList []Node

func (nodes NodeList) Children() []Node {
	return nodes
}

func (nodes NodeList) String() string {
	var b bytes.Buffer

	first := true
	for _, node := range(nodes) {
		if !first {
			b.WriteString(" ")
		}

		b.WriteString(node.String())

		first = false
	}

	return b.String()
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
