#include <assert.h>

#include "item.h"
#include "set.h"
#include "token.h"

const int START_SYMBOL = -1;

int comp_item(const void *a, const void *b) {
	const Item *i1, *i2;
	i1 = a; i2 = b;

	if (i1->rule != i2->rule) {
		return i1->rule - i2->rule;
	}
	
	return i1->pos - i2->pos;
}

int hash_item(const void *item) {
	const Item *i = item;
	return hash(hash(hash_init, i->pos), (intptr_t) i->rule); 
}

Item prod_item(production *prod, int pos) {
	assert(pos <= prod->len);

	Item item;
	item.rule = prod;
	item.pos = pos;

	return item;
}

void write_item(FILE *out, const Grammar *g, const Item item) {
	if (item.rule->head == START_SYMBOL) {
		fprintf(out, "S");
	} else {
		write_token(out, g, item.rule->head);
	}

	fprintf(out, " =>");

	int i;
	for (i = 0; i < item.rule->len; ++i) {
		if (i == item.pos) {
			fprintf(out, " .");
		}

		fprintf(out, " ");
		write_token(out, g, item.rule->tail[i]);
	}

	if (i == item.pos) {
		fprintf(out, " .");
	}

	fprintf(out, "\n");
}
