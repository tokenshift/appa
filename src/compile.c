#include <assert.h>

#include "appa.h"
#include "grammar.h"
#include "item.h"
#include "kernel.h"
#include "map.h"
#include "production.h"

struct Parser {
};

typedef struct {
	Token tkn;
	Kernel kernel;
} k_goto;

int comp_goto(const void *a, const void *b) {
	return ((k_goto *) a)->tkn - ((k_goto *) b)->tkn;
}

int hash_goto(const void *gto) {
	return ((k_goto *) gto)->tkn;
}

Item create_start_item(const Grammar *g, NonTerminal start);
void compute_gotos(const Grammar *g, Kernel *start_kernel);

Parser *appa_compile(const Grammar *g, NonTerminal start) {
	Item start_item = create_start_item(g, start);

	Kernel start_kernel;
	start_kernel.items = create_set(1, sizeof(Item), hash_item, comp_item);
	set_put(start_kernel.items, &start_item);
	
	write_kernel(stdout, g, &start_kernel);
	printf("\n");

	compute_gotos(g, &start_kernel);

	return 0;
}

void compute_gotos(const Grammar *g, Kernel *start_kernel) {
	Kernel closure;
	closure.items = compute_closure(g, start_kernel->items);

	Map *gotos = create_map(8, sizeof(Kernel));
	int i;
	for(i = 0; i < set_len(closure.items); ++i) {
		Item item = *(Item *)set_at(closure.items, i);
		if (item.pos < item.rule->len) {
			Token next = item.rule->tail[item.pos];
			
			Item inext;
			inext.rule = item.rule;
			inext.pos = item.pos + 1;
			inext.lookahead = item.lookahead;

			if (!map_contains(gotos, next)) {
				Kernel k;
				k.items = create_set(4, sizeof(Item), hash_item, comp_item);
				map_put(gotos, next, &k);
			}

			Kernel *kernel = map_get(gotos, next);
			assert(kernel != 0);
			set_put(kernel->items, &inext);
		}
	}

	for (i = 0; i < vec_len(g->tokens); ++i) {
		if (map_contains(gotos, i)) {
			write_kernel(stdout, g, map_get(gotos, i));
			printf("\n");
		}
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
