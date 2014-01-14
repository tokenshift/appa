#include <assert.h>
#include <string.h>

#include "map.h"
#include "vector.h"

struct Map {
	size_t size;
	Vector **data;
};

typedef struct {
	int key;
	void *val;
} map_entry;

Map *map_new(size_t size) {
	assert(size > 0);

	Map *map = calloc(1, sizeof(Map));
	map->size = size;
	map->data = calloc(size, sizeof(Vector *));

	int i;
	for (i = 0; i < size; ++i) {
		map->data[i] = vec_new(sizeof(map_entry), 1);
	}

	return map;
}

void map_delete(Map *m) {
	int i;
	for (i = 0; i < m->size; ++i) {
		vec_delete(m->data[i]);
	}
	free(m);
}

int map_contains(const Map *m, int key) {
	return map_get(m, key) != 0;
}

void *map_get(const Map *m, int key) {
	Vector *v = m->data[key % m->size];

	int i;
	for (i = 0; i < vec_len(v); ++i) {
		map_entry *entry = vec_at(v, i);
		if (entry->key == key) {
			return entry->val;
		}
	}

	return 0;
}

void map_put(Map *m, int key, void *val) {
	Vector *v = m->data[key % m->size];

	int i;
	for (i = 0; i < vec_len(v); ++i) {
		map_entry *entry = vec_at(v, i);
		if (entry->key == key) {
			// Overwrite this entry.
			entry->val = val;
			return;
		}
	}

	// Otherwise, create a new entry.
	map_entry *entry = vec_push(v);
	entry->key = key;
	entry->val = val;
}
