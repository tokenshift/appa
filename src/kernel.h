#ifndef KERNEL_H
#define KERNEL_H

#include <stdio.h>

#include "appa.h"
#include "vector.h"

typedef struct {
	Vector *items;
} Kernel;

void compute_closure(const Grammar *g, const Kernel *kernel, Kernel *closure);
void write_kernel(FILE *out, const Grammar *g, const Kernel *kernel);

#endif
