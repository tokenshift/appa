#ifndef VECTOR_H
#define VECTOR_H

#include <stdlib.h>

typedef struct Vector Vector;

Vector *create_vector(size_t width, int init_length);
void *vec_at(const Vector *v, int index);
void *vec_push(Vector *v);
int vec_len(const Vector *v);

#endif
