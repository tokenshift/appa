#include <assert.h>
#include <stdio.h>

#include "item.h"
#include "grammar.h"
#include "kernel.h"
#include "production.h"
#include "set.h"
#include "token.h"

int hash_int(const void *val) {
	return *(int *)val;
}

int comp_int(const void *a, const void *b) {
	return *(int *)a - *(int *)b;
}

void compute_token_firsts(const Grammar *g, Token t, Set *firsts) {
	Set *processing = create_set(4, sizeof(Token), hash_int, comp_int);
	Set *processed = create_set(4, sizeof(Token), hash_int, comp_int);

	set_put(processing, &t);

	while(!set_empty(processing)) {
		t = *(Token *)set_first(processing);
		set_pop(processing);
		set_put(processed, &t);

		token *tkn = token_info(g, t);
		if (tkn->type == TKN_NONTERM) {
			// Process first token from each production.
			int i;
			for (i = 0; i < vec_len(g->productions); ++i) {
				production *prod = vec_at(g->productions, i);
				if (prod->head == t && prod->len > 0) {
					set_put(processing, &prod->tail[0]);
				}
			}
		} else {
			// Add the terminal token itself.
			if (set_put(processed, &t) == &t) {
				set_put(firsts, &t);
			}
		}
	}

	delete_set(processing);
	delete_set(processed);
}

void compute_closure_items_for_production(const Grammar *g, Item item, production *prod, Set *items, Set *new_items) {
	assert(item.pos + 1 <= item.rule->len);

	// Determine NEXT token.
	Token next;
	if (item.pos + 1 >= item.rule->len) {
		// Use the lookahead
		next = item.lookahead;
	} else {
		next = item.rule->tail[item.pos +1];
	}

	// Compute FIRST terminals for NEXT token.
	Set *firsts = create_set(4, sizeof(Token), hash_int, comp_int);
	compute_token_firsts(g, next, firsts);

	// Create item[0] for the production.
	Item item0;
	item0.rule = prod;
	item0.pos = 0;
	item0.lookahead = next; // TODO
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
			compute_closure_items_for_production(g, item, prod, items, new_items);
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
