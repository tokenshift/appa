package appa

import "bytes"

type kernel struct {
	items []item
}

func createKernel(items ...item) kernel {
	k := kernel {
		make([]item, 0, len(items)),
	}

	for _, it := range(items) {
		k.add(it)
	}

	return k
}

func (k *kernel) add(it item) bool {
	if k.contains(it) {
		return false
	} else {
		k.items = append(k.items, it)
		return true
	}
}

func (k kernel) contains(it item) bool {
	for _, it2 := range(k.items) {
		if it.eq(it2) {
			return true
		}
	}
	return false
}

func (k kernel) String() string {
	var buf bytes.Buffer
	for _, it := range(k.items) {
		buf.WriteString(it.String())
		buf.WriteString("\n")
	}
	return buf.String()
}
