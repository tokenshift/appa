#include <assert.h>
#include <string.h>

#include "map.h"
#include "vector.h"

struct Map {
	Vector *data;
};

typedef struct {
	int key;
	void *val;
} map_entry;

Map *map_new() {
	Map *map = calloc(1, sizeof(Map));
	map->data = vec_new(sizeof(map_entry), 4);
	return map;
}

void map_delete(Map *m) {
	vec_delete(m->data);
	free(m);
}

int map_contains(const Map *m, int key) {
	return map_get(m, key) != 0;
}

map_entry *get_map_entry(const Map *m, int key) {
	int i;
	for (i = 0; i < vec_len(m->data); ++i) {
		map_entry *entry = vec_at(m->data, i);
		if (entry->key == key) {
			return entry;
		}
	}

	return 0;
}

void *map_get(const Map *m, int key) {
	map_entry *entry = get_map_entry(m, key);
	return entry ? entry->val : 0;
}

void map_put(Map *m, int key, void *val) {
	map_entry *entry = get_map_entry(m, key);
	if (entry == 0) {
		entry = vec_push(m->data);
		entry->key = key;
		entry->val = val;
	} else {
		entry->val = val;
	}
}
