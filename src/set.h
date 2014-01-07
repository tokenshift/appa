#ifndef SET_H
#define SET_H

#include <stdlib.h>

typedef struct Set Set;

typedef int (*set_comp_fun)(const void *a, const void *b);
typedef int (*set_hash_fun)(const void *item);

Set *create_set(size_t capacity, set_hash_fun hash, set_comp_fun comp);
void delete_set(Set *set);

// Returns a pointer to the item in the set, or 0 if it is not found.
void *set_find(Set *set, const void *item);

// Checks whether the item is present.
int set_has(Set *set, const void *item);

// Adds an item to the set if it is not already present.
// Returns a pointer to the existing or added item.
void *set_put(Set *set, void *item);

#endif
