package appa

// A set of LALR items.
type lalrSet struct {
	items map[uint32][]lalrItem
}

// Checks whether this set contains the specified LALR item.
func (set lalrSet) contains(item lalrItem) bool {
	hash := item.hash()

	if items, ok := set.items[hash]; ok {
		for _, item2 := range(items) {
			if item.equals(item2) {
				return true
			}
		}
	}

	return false
}

// Value equality for LALR sets.
func (set lalrSet) equals(other lalrSet) bool {
	if set.size() != other.size() {
		return false
	}

	for _, list := range(set.items) {
		for _, item := range(list) {
			if !other.contains(item) {
				return false
			}
		}
	}

	return true
}

// Hashing function for LALR set lookup.
func (set lalrSet) hash() (val uint32) {
	// Need hash to be order independent, so
	// just XOR everything.
	for _, list := range(set.items) {
		for _, item := range(list) {
			val = val ^ item.hash()
		}
	}

	return
}

// The number of items in the set.
func (set lalrSet) size() (count int) {
	for _, list := range(set.items) {
		count = count + len(list)
	}
	return
}
