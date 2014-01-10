#ifndef MAP_H
#define MAP_H

typedef struct Map Map;

Map *create_map(size_t size, size_t width);
void delete_map(Map *m);

int map_contains(const Map *m, int key);
void *map_get(const Map *m, int key);
void map_put(Map *m, int key, const void *val);

#endif
