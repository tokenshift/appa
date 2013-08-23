package appa

type lit struct {
	text string
}

func (l *lit) String() string {
	return l.text
}
