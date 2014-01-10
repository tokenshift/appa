#include <assert.h>
#include <string.h>

#include "vector.h"

struct Vector {
	// The size of elements contained by the vector.
	size_t width;

	// The total allocated size (in bytes) of the underlying array.
	size_t size;

	// The byte offset of the used portion of the underlying array.
	int offset;

	// The size (in bytes) of the used portion of the array.
	int length;

	// The underlying array containing the data.
	void *data;
};

Vector *create_vector(size_t width, int init_length) {
	assert(width > 0);
	assert(init_length >= 0);

	Vector *v = calloc(1, sizeof(Vector));
	v->size = width * init_length;
	v->width = width;
	v->data = calloc(init_length, width);

	return v;
}

// Destroys the vector, freeing all allocated memory.
void delete_vector(Vector *v) {
	free(v->data);
	free(v);
}

// Returns a pointer to the element at the specified index.
void *vec_at(const Vector *v, int index) {
	assert((index * v->width) < v->length);
	return v->data + v->offset + index*v->width;
}

// Expands the underlying array of the vector.
void vec_expand(Vector *v) {
	size_t size = 1.5 * v->size + 1;
	void *data = malloc(size);
	memcpy(data, v->data + v->offset, v->length);
	memset(data + v->length, 0, size - v->length);
	v->size = size;
	v->offset = 0;

	free(v->data);
	v->data = data;
}

// Removes the first item from the vector.
void vec_pop(Vector *v) {
	assert(vec_len(v) > 0);
	v->offset += v->width;
	v->length -= v->width;
}

// Returns a pointer to a chunk of memory appended to the vector.
void *vec_push(Vector *v) {
	while (v->offset + v->length + v->width >= v->size) {
		vec_expand(v);
	}

	v->length += v->width;
	return v->data + v->offset + v->length - v->width;
}

// Returns the number of elements in the vector.
int vec_len(const Vector *v) {
	return v->length / v->width;
}
