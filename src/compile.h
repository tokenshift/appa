#ifndef COMPILE_H
#define COMPILE_H

#include "appa.h"
#include "item.h"
#include "kernel.h"
#include "kernel_set.h"

Item create_start_item(const Grammar *g, NonTerminal start);
KernelSet *compute_gotos(Kernel *start_kernel);

#endif
