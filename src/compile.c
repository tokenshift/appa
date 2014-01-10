#include "appa.h"
#include "grammar.h"
#include "item.h"
#include "kernel.h"
#include "production.h"

struct Parser {
};

Item create_start_item(const Grammar *g, NonTerminal start);

Parser *appa_compile(const Grammar *g, NonTerminal start) {
	Item start_item = create_start_item(g, start);

	Kernel start_kernel;
	start_kernel.items = create_set(1, sizeof(Item), hash_item, comp_item);
	set_put(start_kernel.items, &start_item);
	
	write_kernel(stdout, g, &start_kernel);
	printf("\n");

	Kernel closure;
	closure.items = compute_closure(g, start_kernel.items);

	write_kernel(stdout, g, &closure);

	return 0;
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
