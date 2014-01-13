#include <assert.h>

#include "appa.h"
#include "grammar.h"
#include "item.h"
#include "kernel.h"
#include "map.h"
#include "production.h"

struct Parser {
};

int comp_kernel(const void *a, const void *b) {
	const Kernel *k1, *k2;
	k1 = a;
	k2 = b;

	if (set_len(k1->items) != set_len(k2->items)) {
		return set_len(k1->items) - set_len(k2->items);
	}

	int i;
	for (i = 0; i < set_len(k1->items); ++i) {
		Item *item = set_at(k1->items, i);
		if (set_find(k2->items, item) == 0) {
			return -1;
		}
	}

	return 0;
}

int hash_kernel(const void *k) {
	const Kernel *kernel = k;
	int i;
	int hash = 0;
	for (i = 0; i < set_len(kernel->items); ++i) {
		hash = hash ^ hash_item(set_at(kernel->items, i));
	}

	return hash;
}

Item create_start_item(const Grammar *g, NonTerminal start);
void compute_gotos(const Grammar *g, Kernel *start_kernel);

Parser *appa_compile(const Grammar *g, NonTerminal start) {
	Item start_item = create_start_item(g, start);

	Kernel start_kernel;
	start_kernel.items = create_set(1, sizeof(Item), hash_item, comp_item);
	start_kernel.gotos = create_map(4, sizeof(Kernel *));
	set_put(start_kernel.items, &start_item);
	
	write_kernel(stdout, g, &start_kernel);
	printf("\n");

	compute_gotos(g, &start_kernel);

	return 0;
}

void compute_gotos(const Grammar *g, Kernel *start_kernel) {
	Set *kernels = create_set(8, sizeof(Kernel), hash_kernel, comp_kernel);
	Set *new_kernels = create_set(8, sizeof(Kernel), hash_kernel, comp_kernel);
	set_put(kernels, start_kernel);
	set_put(new_kernels, start_kernel);

	Kernel *kernel;
	while ((kernel = set_first(new_kernels)) != 0) {
		Kernel closure;
		closure.items = compute_closure(g, kernel->items);
		closure.gotos = 0;

		Item *item;
		while ((item = set_first(closure.items)) != 0) {
			if (item->pos < item->rule->len) {
				Token next = item->rule->tail[item->pos];

				Item inext;
				inext.rule = item->rule;
				inext.pos = item->pos + 1;
				inext.lookahead = item->lookahead;

				Kernel *kgoto;
				if ((kgoto = map_get(kernel->gotos, next)) == 0) {
					kgoto = calloc(1, sizeof(Kernel));
					kgoto->items = create_set(4, sizeof(Item), hash_item, comp_item);
					kgoto->gotos = create_map(4, sizeof(Kernel *));
					map_put(kernel->gotos, next, kgoto);
				}

				set_put(kgoto->items, &inext);
			}

			set_pop(closure.items);
		}

		int i;
		for (i = 0; i < vec_len(g->tokens); ++i) {
			if (map_contains(kernel->gotos, i)) {
				write_kernel(stdout, g, map_get(kernel->gotos, i));
				printf("\n");
			}
		}

		delete_set(closure.items);
		set_pop(new_kernels);
	}
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
