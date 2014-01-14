#include <assert.h>

#include "item.h"
#include "token.h"

int item_core_eq(Item a, Item b) {
	return a.rule == b.rule && a.pos == b.pos;
}

int item_eq(Item a, Item b) {
	return item_core_eq(a, b) && a.lookahead == b.lookahead;
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
}
