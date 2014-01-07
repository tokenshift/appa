#include "appa.h"
#include "item.h"
#include "kernel.h"
#include "production.h"
#include "vector.h"

struct Parser {
};

void create_start_item(const Grammar *g, NonTerminal start, Item *item);

Parser *appa_compile(const Grammar *g, NonTerminal start) {
	Kernel start_kernel;
	start_kernel.items = create_vector(sizeof(Item), 1);
	
	Item *start_item = vec_push(start_kernel.items);
	create_start_item(g, start, start_item);

	//write_kernel(stdout, g, &start_kernel);

	Kernel closure;
	compute_closure(g, &start_kernel, &closure);
	write_kernel(stdout, g, &closure);

	return 0;
}

void create_start_item(const Grammar *g, NonTerminal start, Item *item) {
	item->rule = calloc(1, sizeof(Production));
	item->rule->head = START_SYMBOL;
	item->rule->len = 1;
	item->rule->tail = calloc(1, sizeof(Token));
	item->rule->tail[0] = start;
	item->pos = 0;
}
