#ifndef MAP_H
#define MAP_H

// Collection mapping integers to pointers.
typedef struct Map Map;

// Creates a new map.
Map *map_new(size_t size);

// Deletes a map.
void map_delete(Map *m);

// Checks whether the specified key is found in the map.
int map_contains(const Map *m, int key);

// Gets the item with the specified key.
void *map_get(const Map *m, int key);

// Adds an item with the specified key.
void map_put(Map *m, int key, void *val);

#endif
