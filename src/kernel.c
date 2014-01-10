#include <stdio.h>

#include "item.h"
#include "grammar.h"
#include "kernel.h"
#include "production.h"
#include "set.h"
#include "token.h"

void compute_closure_items_for_production(Item item, production *prod, Set *items, Set *new_items) {
	// Create item[0] for the production.
	Item item0;
	item0.rule = prod;
	item0.pos = 0;
	item0.lookahead = EOF_SYMBOL; // TODO
	// Put item[0] in kernel
	if (set_put(items, &item0) == &item0) {
		// If it is new, put item[0] in new_item_set
		set_put(new_items, &item0);
	}
}

void compute_closure_items(const Grammar *g, Item item, Set *items, Set *new_items) {
	if (item.pos >= item.rule->len) return;
	
	// If FIRST(item) is a non-terminal:
	Token t = item.rule->tail[item.pos];
	token *tkn = token_info(g, t);
	if (tkn->type != TKN_NONTERM) return;

	// For each production for that non-terminal:
	int i;
	for (i = 0; i < vec_len(g->productions); ++i) {
		production *prod = vec_at(g->productions, i);
		if (prod->head == t) {
			compute_closure_items_for_production(item, prod, items, new_items);
		}
	}
}

Set *compute_closure(const Grammar *g, const Set *kernel) {
	Set *new_items = create_set(4, sizeof(Item), hash_item, comp_item);
	Set *items = create_set(4, sizeof(Item), hash_item, comp_item);

	Item *item;
	
	// Put all kernel items in the set of new items.
	int i;
	for (i = 0; i < set_len(kernel); ++i) {
		item = set_at(kernel, i);
		set_put(new_items, item);
		set_put(items, item);
	}

	// For each item in the new item set:
	while((item = set_first(new_items)) != 0) {
		Item it = *item;
		set_pop(new_items);
		compute_closure_items(g, it, items, new_items);
	}

	delete_set(new_items);
	return items;
}

void write_kernel(FILE *out, const Grammar *g, const Kernel *kernel) {
	int i;
	for (i = 0; i < set_len(kernel->items); ++i) {
		Item *item = set_at(kernel->items, i);
		write_item(out, g, *item);
	}
}
