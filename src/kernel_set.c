#include "kernel_set.h"
#include "vector.h"

struct KernelSet {
	Vector *kernels;
};

KernelSet *kernel_set_new() {
	KernelSet *set = calloc(1, sizeof(KernelSet));
	set->kernels = vec_new(sizeof(Kernel *), 4);
	return set;
}

void kernel_set_delete(KernelSet *set) {
	vec_delete(set->kernels);
	free(set);
}

void kernel_set_add(KernelSet *set, Kernel *k) {
	*(Kernel **)vec_push(set->kernels) = k;
}

Kernel *kernel_set_at(const KernelSet *set, int index) {
	return *(Kernel **)vec_at(set->kernels, index);
}

int kernel_set_len(const KernelSet *set) {
	return vec_len(set->kernels);
}

Kernel *kernel_set_find_by_core(const KernelSet *set, Kernel *kernel) {
	int i;
	for (i = 0; i < kernel_set_len(set); ++i) {
		Kernel *candidate = kernel_set_at(set, i);
		if (kernel_core_eq(kernel, candidate)) {
			return candidate;
		}
	}

	return 0;
}
