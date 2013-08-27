package appa

import "bytes"
import "fmt"
import "hash/fnv"
import "io"

type rule struct {
	body []Token
}

func createRule(tkns ...Token) rule {
	return rule { tkns }
}

func (r rule) at(i int) Token {
	return r.body[i]
}

func (r rule) equals(r2 rule) bool {
	if len(r.body) != len(r2.body) {
		return false
	}

	for i, tkn := range(r.body) {
		if tkn != r2.body[i] {
			return false
		}
	}

	return true
}

// Hash function for LALR item lookup.
func (r rule) hash() uint32 {
	hash := fnv.New32()

	for _, tkn := range(r.body) {
		io.WriteString(hash, fmt.Sprint(tkn))
	}

	return hash.Sum32()
}

func (r rule) size() int {
	return len(r.body)
}

func (r rule) String() string {
	out := new(bytes.Buffer)

	first := true
	for _, tkn := range(r.body) {
		if !first {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, tkn)
		first = false
	}

	return out.String()
}
