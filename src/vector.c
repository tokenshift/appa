#include <assert.h>
#include <string.h>

#include "vector.h"

struct Vector {
	// The size of the underlying array (number of elements).
	size_t size;

	// The offset of the used portion of the underlying array.
	int offset;

	// The number of elements in the used portion of the array.
	int length;

	// The underlying array containing the data.
	void **data;
};

Vector *create_vector(size_t size) {
	Vector *v = calloc(1, sizeof(Vector));
	v->size = size;
	v->data = calloc(size, sizeof(void *));
	return v;
}

void delete_vector(Vector *v) {
	free(v->data);
	free(v);
}

void *vec_at(const Vector *v, int index) {
	assert(index < v->length);
	return v->data[index];
}

void vec_expand(Vector *v) {
	size_t size = 1.5 * v->size + 1;
	void **data = calloc(size, sizeof(void *));

	memcpy(data, &v->data[v->offset], v->length * sizeof(void *));
	memset(&data[v->length], 0, (size - v->length) * sizeof(void *));
	v->size = size;
	v->offset = 0;

	free(v->data);
	v->data = data;
}

void *vec_pop(Vector *v) {
	assert(v->length > 0);
	
	void *item = v->data[0];
	++v->offset;
	--v->length;
	return item;
}

void vec_push(Vector *v, void *item) {
	if (v->offset + v->length >= v->size) {
		vec_expand(v);
	}

	++v->length;
	v->data[v->offset + v->length] = item;
}

int vec_len(const Vector *v) {
	return v->length;
}
