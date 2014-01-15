#ifndef KERNEL_SET_H
#define KERNEL_SET_H

#include "kernel.h"

typedef struct KernelSet KernelSet;

// Creates a new kernel set.
KernelSet *kernel_set_new();

// Deletes a kernel set.
void kernel_set_delete(KernelSet *set);

// Adds a kernel to the kernel set.
void kernel_set_add(KernelSet *set, Kernel *k);

// Gets the kernel at the specified index.
Kernel *kernel_set_at(const KernelSet *set, int index);

// Gets the number of kernels in the kernel set.
int kernel_set_len(const KernelSet *set);

// Finds a kernel in the set whose core (items without lookaheads) matches the
// core of the specified kernel.
Kernel *kernel_set_find_by_core(const KernelSet *set, Kernel *k);

#endif
