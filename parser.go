package appa

type Parser interface {
	ParseString(input string) (result Node, err error)
}
