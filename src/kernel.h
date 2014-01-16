#ifndef KERNEL_H
#define KERNEL_H

#include <stdio.h>

#include "appa.h"
#include "item_set.h"
#include "map.h"

typedef struct {
	const Grammar *g;
	ItemSet *items;
	Map *gotos;
} Kernel;

// Creates a new LALR kernel.
Kernel *kernel_new(const Grammar *g);

// Deletes an LALR kernel.
void kernel_delete(Kernel *k);

// Adds an item to the kernel. Returns 1 if the item was added; 0 if it already
// existed.
int kernel_add(Kernel *k, Item item);

// Gets the item at the specified index.
Item kernel_at(const Kernel *k, int index);

// Computes the closure of the kernel.
ItemSet *kernel_closure(const Kernel *k);

// Checks whether the cores of the kernels (items without lookaheads) match.
int kernel_core_eq(const Kernel *a, const Kernel *b);

// Gets the number of items in the kernel.
int kernel_len(const Kernel *k);

// Writes the kernel contents (for debugging).
void write_kernel(const Kernel *k, FILE *out);

#endif
