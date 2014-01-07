#include <stdio.h>

#include "item.h"
#include "grammar.h"
#include "kernel.h"
#include "production.h"
#include "set.h"

void compute_closure(const Grammar *g, const Kernel *kernel, Kernel *closure) {
	Set *new_item_set = create_set(4, hash_item, comp_item);
	Set *items = create_set(4, hash_item, comp_item);
	closure->items = create_vector(sizeof(Item), vec_len(kernel->items));

	Item *item;
	
	// Put all kernel items in the set of new items.
	int i;
	for (i = 0; i < vec_len(kernel->items); ++i) {
		Item *item = vec_at(kernel->items, i);
		set_put(new_item_set, item);

		Item *it = vec_push(closure->items);
		it->rule = item->rule;
		it->pos = item->pos;
	}

	// For each item in the new item set:
	while((item = set_pop(new_item_set)) != 0) {
		if (item->pos < item->rule->len) {
			// If FIRST(item) is a non-terminal:
			Token t = item->rule->tail[item->pos];
			tkn_info *tkn = token_at(g, t);
			if (tkn->type == TKN_NONTERM) {
				// For each production for that non-terminal:
				int i;
				for (i = 0; i < vec_len(g->productions); ++i) {
					if (((Production *)vec_at(g->productions, i))->head == t) {	
						// Create item[0] from that production
						Item *item0 = prod_item(vec_at(g->productions, i), 0);
						// Put item[0] in kernel
						if (set_put(items, item0) == item0) {
							// If it is new, put item[0] in new_item_set
							set_put(new_item_set, item0);
						}
					}
				}
			}
		}
	}

	while((item = set_pop(items)) != 0) {
		Item *it = vec_push(closure->items);
		// TODO: Figure out a way to do this without allocating/deallocating
		// the items.
		it->rule = item->rule;
		it->pos = item->pos;
		free(item);
	}

	delete_set(new_item_set);
	delete_set(items);
}

void write_kernel(FILE *out, const Grammar *g, const Kernel *kernel) {
	int i;
	for (i = 0; i < vec_len(kernel->items); ++i) {
		Item *item = vec_at(kernel->items, i);
		write_item(out, g, item);
	}
}
