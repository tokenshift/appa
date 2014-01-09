#include <assert.h>
#include <string.h>

#include "set.h"
#include "vector.h"

struct Set {
	size_t size;
	size_t width;
	set_hash_fun hash;
	set_comp_fun comp;

	Vector **items;
};

Set *create_set(size_t size, size_t width, set_hash_fun hash, set_comp_fun comp) {
	assert(size > 0);
	assert(width > 0);

	Set *set = malloc(sizeof(Set) + size*sizeof(Vector *));
	set->size = size;
	set->width = width;
	set->hash = hash;
	set->comp = comp;
	set->items = (Vector **)(set + sizeof(Set));

	int i;
	for (i = 0; i < size; ++i) {
		set->items[i] = create_vector(width, 1);
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

void *set_alter(Set *set, const void *find, void *put) {
	int hash = set->hash(find);

	Vector *v = set->items[hash % set->size];
	int i;
	for (i = 0; i < vec_len(v); ++i) {
		void *item = vec_at(v, i);
		if (set->comp(item, find) == 0) {
			return put ? 0 : item;
		}
	}

	if (put != 0) {
		void *space = vec_push(v);
		memcpy(space, put, set->width);
	}

	return 0;
}

void *set_put(Set *set, void *item) {
	return set_alter(set, item, item);
}

void *set_find(Set *set, const void *item) {
	return set_alter(set, item, 0);
}

int set_has(Set *set, const void *item) {
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

void *set_first(Set *set) {
	return set_pop_first(set, 0);
}

void set_pop(Set *set) {
	set_pop_first(set, 1);
}

const int hash_init = 2166136261;
int hash(int hash, int val) {
	return (hash * 16777619) ^ val;
}
