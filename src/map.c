#include <assert.h>
#include <string.h>

#include "map.h"
#include "vector.h"

struct Map {
	size_t size;
	size_t width;

	Vector **data;
};

Map *create_map(size_t size, size_t width) {
	assert(size > 0);
	assert(width > 0);

	Map *map = calloc(1, sizeof(Map));
	map->size = size;
	map->width = width;
	map->data = calloc(size, sizeof(Vector *));

	int i;
	for (i = 0; i < size; ++i) {
		map->data[i] = create_vector(sizeof(int) + width, 1);
	}

	return map;
}

void delete_map(Map *m) {
	int i;
	for (i = 0; i < m->size; ++i) {
		delete_vector(m->data[i]);
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
		int *entry = vec_at(v, i);
		if (*entry == key) {
			return entry + sizeof(int);
		}
	}

	return 0;
}

void map_put(Map *m, int key, const void *val) {
	Vector *v = m->data[key % m->size];

	int i, *entry;
	for (i = 0; i < vec_len(v); ++i) {
		entry = vec_at(v, i);
		if (*entry == key) {
			// Overwrite this entry.
			memcpy(entry + sizeof(int), val, m->width);
			return;
		}
	}

	// Otherwise, create a new entry.
	entry = vec_push(v);
	*entry = key;
	memcpy(entry + sizeof(int), val, m->width);
}
