#ifndef VECTOR_H
#define VECTOR_H

#include <stdlib.h>

typedef struct Vector Vector;

// Creates a new vector with the specified initial capacity.
Vector *create_vector(size_t size);

// Frees memory allocated for a vector.
void delete_vector(Vector *v);

// Returns the item at the specified index.
void *vec_at(const Vector *v, int index);

// Removes and returns the first item in the vector.
void *vec_pop(Vector *v);

// Appends an item to the vector.
void vec_push(Vector *v, void *item);

// Gets the length of the vector.
int vec_len(const Vector *v);

#endif
