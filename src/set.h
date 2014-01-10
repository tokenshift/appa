#ifndef SET_H
#define SET_H

#include <stdlib.h>

typedef struct Set Set;

typedef int (*set_comp_fun)(const void *a, const void *b);
typedef int (*set_hash_fun)(const void *item);

Set *create_set(size_t capacity, size_t width, set_hash_fun hash, set_comp_fun comp);
void delete_set(Set *set);

// Returns a pointer to the item in the set, or 0 if it is not found.
void *set_find(Set *set, const void *item);

// Checks whether the set is empty.
int set_empty(const Set *set);

// Checks whether the item is present.
int set_has(Set *set, const void *item);

// Returns a pointer to the first item in the set.
void *set_first(Set *set);

// Deletes the first item in the set.
void set_pop(Set *set);

// Copies an item into the set, if it does not already exist.
void *set_put(Set *set, void *item);

// Returns the length of the set.
int set_len(const Set *set);

// Returns a pointer to the specified element.
void *set_at(const Set *set, int index);

// Simple utility hash function.
int hash(int hash, intptr_t val);
extern const int hash_init;

#endif
