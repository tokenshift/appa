#include <assert.h>

#include "item.h"
#include "set.h"
#include "token.h"

int comp_item(const void *a, const void *b) {
	const Item *i1, *i2;
	i1 = a; i2 = b;

	if (i1->rule != i2->rule) {
		return i1->rule - i2->rule;
	}

	if (i1->lookahead != i2->lookahead) {
		return i1->lookahead - i2->lookahead;
	}
	
	return i1->pos - i2->pos;
}

int hash_item(const void *item) {
	const Item *i = item;
	return hash(hash(hash(hash_init, i->pos), (intptr_t) i->rule), i->lookahead);
}

void write_item(FILE *out, const Grammar *g, const Item item) {
	write_token(out, g, item.rule->head);
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

	fprintf(out, ", ");
	write_token(out, g, item.lookahead);
	fprintf(out, "\n");
}
