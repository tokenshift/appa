#ifndef KERNEL_H
#define KERNEL_H

#include <stdio.h>

#include "appa.h"
#include "set.h"

typedef struct {
	Set *items;
} Kernel;

Set *compute_closure(const Grammar *g, const Set *kernel);
void write_kernel(FILE *out, const Grammar *g, const Kernel *kernel);

int hash_int(const void *val);
int comp_int(const void *a, const void *b);

#endif
