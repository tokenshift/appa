#include <assert.h>
#include <string.h>

#include "set.h"
#include "vector.h"

struct Set {
	size_t size;
	set_hash_fun hash;
	set_comp_fun comp;

	Vector **items;
};

Set *create_set(size_t size, set_hash_fun hash, set_comp_fun comp) {
	assert(size > 0);

	Set *set = calloc(1, sizeof(Set));
	set->size = size;
	set->hash = hash;
	set->comp = comp;
	set->items = calloc(size, sizeof(Vector *));

	int i;
	for (i = 0; i < size; ++i) {
		set->items[i] = create_vector(1);
	}

	return set;
}

void delete_set(Set *set) {
	int i;
	for (i = 0; i < set->size; ++i) {
		delete_vector(set->items[i]);
	}

	free(set);
}

void *set_put(Set *set, void *item) {
	void *found = set_find(set, item);
	if (found) return found;

	int hash = set->hash(item);
	Vector *v = set->items[hash % set->size];
	vec_push(v, item);
	return 0;
}

void *set_find(const Set *set, const void *item) {
	int hash = set->hash(item);

	Vector *v = set->items[hash % set->size];
	int i;
	for (i = 0; i < vec_len(v); ++i) {
		void *it = vec_at(v, i);
		if (set->comp(it, item) == 0) {
			return it;
		}
	}

	return 0;
}

int set_has(const Set *set, const void *item) {
	return set_find(set, item) != 0;
}

void *set_pop_first(Set *set, int delete) {
	int i;
	for (i = 0; i < set->size; ++i) {
		Vector *v = set->items[i];
		if (vec_len(v) > 0) {
			if (delete) {
				vec_pop(v);
				return 0;
			} else {
				return vec_at(v, 0);
			}
		}
	}

	return 0;
}

void *set_first(const Set *set) {
	int i;
	for (i = 0; i < set->size; ++i) {
		Vector *v = set->items[i];
		if (vec_len(v) > 0) {
			return vec_at(v, 0);
		}
	}

	return 0;
}

void *set_pop(Set *set) {
	int i;
	for (i = 0; i < set->size; ++i) {
		Vector *v = set->items[i];
		if (vec_len(v) > 0) {
			return vec_pop(v);
		}
	}

	return 0;
}

// Returns the length of the set.
int set_len(const Set *set) {
	int i, count;
	for (i = 0, count = 0; i < set->size; ++i) {
		count += vec_len(set->items[i]);
	}
	return count;
}

// Checks whether the set is empty.
int set_empty(const Set *set) {
	return set_len(set) == 0;
}

// Returns a pointer to the specified element.
void *set_at(const Set *set, int index) {
	assert(index < set_len(set));

	int i;
	for (i = 0; i < set->size; ++i) {
		if (index >= vec_len(set->items[i])) {
			index -= vec_len(set->items[i]);
		} else {
			return vec_at(set->items[i], index);
		}
	}

	return 0;
}

const int hash_init = 2166136261;
int hash(int hash, intptr_t val) {
	return (hash * 16777619) ^ val;
}
