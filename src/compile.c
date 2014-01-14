#include <assert.h>

#include "appa.h"
#include "compile.h"
#include "grammar.h"
#include "item.h"
#include "item_set.h"
#include "kernel.h"
#include "kernel_set.h"
#include "map.h"
#include "production.h"

struct Parser {
};

KernelSet *compute_gotos(const Kernel *start_kernel);

Kernel *create_start_kernel(const Grammar *g, NonTerminal start) {
	Kernel *kernel = kernel_new(g);
	item_set_add(kernel->items, create_start_item(g, start));
	return kernel;
}

Parser *appa_compile(const Grammar *g, NonTerminal start) {
	Kernel *start_kernel = create_start_kernel(g, start);
	compute_gotos(start_kernel);

	return 0;
}

void appa_write_dot_grammar(const Grammar *g, FILE *out, NonTerminal start) {
	Kernel *start_kernel = create_start_kernel(g, start);
	KernelSet *kernels = compute_gotos(start_kernel);

	fprintf(out, "digraph {\n\tnode [shape=box];\n");

	int k;
	for (k = 0; k < kernel_set_len(kernels); ++k) {
		const Kernel *kernel = kernel_set_at(kernels, k);

		if (kernel == start_kernel) {
			fprintf(out, "\n\t\"%p\" [style=bold, label=\"", kernel);
		} else {
			fprintf(out, "\n\t\"%p\" [label=\"", kernel);
		}

		int i;
		for (i = 0; i < kernel_len(kernel); ++i) {
			Item item = kernel_at(kernel, i);
			if (i != 0) {
				fprintf(out, "\\n");
			}

			write_item(out, g, item);
		}
		fprintf(out, "\"];\n");

		int t;
		for (t = 0; t < vec_len(g->tokens); ++t) {
			if (map_contains(kernel->gotos, t)) {
				fprintf(out, "\t\"%p\" -> \"%p\";\n", kernel, map_get(kernel->gotos, t));
			}
		}
	}

	fprintf(out, "}\n");
}

KernelSet *compute_gotos(const Kernel *start_kernel) {
	KernelSet *kernels = kernel_set_new();
	kernel_set_add(kernels, start_kernel);

	int k;
	for (k = 0; k < kernel_set_len(kernels); ++k) {
		Kernel *kernel = kernel_set_at(kernels, k);

		// Build the GOTO kernels for this kernel.
		ItemSet *closure = kernel_closure(kernel);

		int i;
		for (i = 0; i < item_set_len(closure); ++i) {
			Item item = item_set_at(closure, i);
			if (item.pos < item.rule->len) {
				Token next = item.rule->tail[item.pos];

				Item inext;
				inext.rule = item.rule;
				inext.pos = item.pos + 1;
				inext.lookahead = item.lookahead;

				Kernel *kgoto;
				if ((kgoto = map_get(kernel->gotos, next)) == 0) {
					kgoto = kernel_new(kernel->g);
					map_put(kernel->gotos, next, kgoto);
				}

				kernel_add(kgoto, inext);
			}
		}

		item_set_delete(closure);

		// Add all of the GOTO kernels to the set of kernels.
		for (i = 0; i < vec_len(kernel->g->tokens); ++i) {
			Kernel *kgoto = map_get(kernel->gotos, i);
			if (kgoto != 0) {
				Kernel *actual = kernel_set_find_by_core(kernels, kgoto);
				if (actual == 0) {
					// If the kernel was not already in the set, add it to the
					// set of kernels to be processed/expanded.
					kernel_set_add(kernels, kgoto);
				} else {
					// Otherwise, delete the duplicate kernel and repoint at
					// the existing one.
					// TODO: Merge lookaheads.
					kernel_delete(kgoto);
					map_put(kernel->gotos, i, actual);
				}
			}
		}
	}

	return kernels;
}

Item create_start_item(const Grammar *g, NonTerminal start) {
	Item item;
	item.rule = calloc(1, sizeof(production));
	item.rule->head = TKN_START;
	item.rule->len = 1;
	item.rule->tail = calloc(1, sizeof(Token));
	item.rule->tail[0] = start;
	item.pos = 0;
	item.lookahead = TKN_EOF;
	return item;
}
