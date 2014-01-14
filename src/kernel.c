#include <assert.h>
#include <stdio.h>

#include "item.h"
#include "grammar.h"
#include "kernel.h"
#include "production.h"
#include "token.h"
#include "token_set.h"

void compute_closure_items(const Grammar *g, Item item, ItemSet *closure);

Kernel *kernel_new(const Grammar *g) {
	Kernel *k = calloc(1, sizeof(Kernel));
	k->g = g;
	k->items = item_set_new();
	k->gotos = map_new(4);
	return k;
}

void kernel_delete(Kernel *k) {
	item_set_delete(k->items);
	map_delete(k->gotos);
	free(k);
}

void kernel_add(Kernel *k, Item item) {
	item_set_add(k->items, item);
}

Item kernel_at(const Kernel *k, int index) {
	return item_set_at(k->items, index);
}

ItemSet *kernel_closure(const Kernel *k) {
	ItemSet *closure = item_set_new();
	
	// Put all kernel items in the set of new items.
	int i;
	for (i = 0; i < kernel_len(k); ++i) {
		Item item = kernel_at(k, i);
		item_set_add(closure, item);
	}

	for (i = 0; i < item_set_len(closure); ++i) {
		compute_closure_items(k->g, item_set_at(closure, i), closure);
	}

	return closure;
}

int max(int a, int b) {
	return a > b ? a : b;
}

int kernel_core_eq(const Kernel *a, const Kernel *b) {
	// Every item in a should be found in b, and vice versa.
	int alen = kernel_len(a);
	int blen = kernel_len(b);

	int ai, bi;
	for (ai = 0; ai < alen; ++ai) {
		Item item = kernel_at(a, ai);
		int match = 0;
		for (bi = 0; bi < blen; ++bi) {
			Item other = kernel_at(b, bi);
			if (item_core_eq(item, other)) {
				match = 1;
				break;
			}
		}

		if (!match) return 0;
	}

	for (bi = 0; bi < blen; ++bi) {
		Item item = kernel_at(b, bi);
		int match = 0;
		for (ai = 0; ai < alen; ++ai) {
			Item other = kernel_at(a, ai);
			if (item_core_eq(item, other)) {
				match = 1;
				break;
			}
		}

		if (!match) return 0;
	}

	return 1;
}

int kernel_len(const Kernel *k) {
	return item_set_len(k->items);
}

void write_kernel(const Kernel *kernel, FILE *out) {
	int i;
	for (i = 0; i < item_set_len(kernel->items); ++i) {
		Item item = item_set_at(kernel->items, i);
		write_item(out, kernel->g, item);
	}
}

TokenSet *compute_token_firsts(const Grammar *g, Token t) {
	TokenSet *firsts = token_set_new();
	token_set_add(firsts, t);

	int i;
	for (i = 0; i < token_set_len(firsts); ++i) {
		Token tkn = token_set_at(firsts, i);
		token *token = token_info(g, tkn);

		if (token->type == TKN_NONTERM) {
			// Process first token from each production.
			int j;
			for (j = 0; j < vec_len(g->productions); ++j) {
				production *prod = vec_at(g->productions, j);
				if (prod->head == t && prod->len > 0) {
					token_set_add(firsts, prod->tail[0]);
				}
			}
		} else {
			// Add the terminal token itself.
			token_set_add(firsts, tkn);
		}
	}

	return firsts;
}

void compute_closure_items_for_production(const Grammar *g, Item item, production *prod, ItemSet *closure) {
	assert(item.pos + 1 <= item.rule->len);

	// Determine NEXT token.
	Token next;
	if (item.pos + 1 >= item.rule->len) {
		// Use the lookahead
		next = item.lookahead;
	} else {
		next = item.rule->tail[item.pos + 1];
	}

	// Compute FIRST terminals for NEXT token.
	TokenSet *firsts = compute_token_firsts(g, next);

	// Create item[0] for the production.
	Item item0;
	item0.rule = prod;
	item0.pos = 0;
	item0.lookahead = next; // TODO: iterate for FIRSTs.
	item_set_add(closure, item0);

	token_set_delete(firsts);
}

void compute_closure_items(const Grammar *g, Item item, ItemSet *closure) {
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
			compute_closure_items_for_production(g, item, prod, closure);
		}
	}
}
