#include <assert.h>
#include <string.h>

#include "vector.h"

struct Vector {
	// The size of the elements in the array.
	size_t width;

	// The size of the underlying array (number of elements).
	size_t size;

	// The offset of the used portion of the underlying array.
	int offset;

	// The number of elements in the used portion of the array.
	int length;

	// The underlying array.
	void *data;
};

void *vec_new(size_t width, size_t capacity) {
	assert(width > 0);
	assert(capacity >= 0);

	Vector *v = calloc(1, sizeof(Vector));
	v->width = width;
	v->size = capacity;
	v->data = calloc(capacity, width);

	return v;
}

void vec_delete(Vector *v) {
	free(v->data);
	free(v);
}

void *vec_at(const Vector *v, int index) {
	assert(index < v->length);
	return v->data + (v->offset + index)*v->width;
}

int vec_len(const Vector *v) {
	return v->length;
}

void expand_vector(Vector *v) {
	size_t size = 1.5*v->size + 1;
	void * used_start = v->data + v->offset*v->width;
	size_t used_length = v->length*v->width;

	void *data = calloc(size, v->width);
	memcpy(data, used_start, used_length);
	memset(data + used_length, 0, size*v->width - used_length);
	free(v->data);

	v->data = data;
	v->size = size;
}

void *vec_push(Vector *v) {
	if (v->length + v->offset >= v->size) {
		// Expand the vector.
		expand_vector(v);
	}

	++v->length;
	return v->data + (v->offset + v->length - 1)*v->width;
}
