#ifndef VECTOR_H
#define VECTOR_H

#include <stdlib.h>

typedef struct Vector Vector;

// Creates a new vector.
void *vec_new(size_t width, size_t capacity);

// Deletes a vector.
void vec_delete(Vector *v);

// Gets the element at the specified index.
void *vec_at(const Vector *v, int index);

// Returns the number of elements in the vector.
int vec_len(const Vector *v);

// Allocates space for a new element at the end of the vector.
void *vec_push(Vector *v);

#endif
