#include <assert.h>

#include "item.h"
#include "token.h"

int item_core_eq(Item a, Item b) {
	return a.rule == b.rule && a.pos == b.pos;
}

int item_eq(Item a, Item b) {
	return item_core_eq(a, b) && a.lookahead == b.lookahead;
}

void write_item_helper(FILE *out, const Grammar *g, const Item item, int escaped) {
	write_token(out, g, item.rule->head);
	fprintf(out, " =>");

	int i;
	for (i = 0; i < item.rule->len; ++i) {
		if (i == item.pos) {
			fprintf(out, " .");
		}

		fprintf(out, " ");
		if (escaped) {
			write_token_escaped(out, g, item.rule->tail[i]);
		} else {
			write_token(out, g, item.rule->tail[i]);
		}
	}

	if (i == item.pos) {
		fprintf(out, " .");
	}

	fprintf(out, ", ");
	if (escaped) {
		write_token_escaped(out, g, item.lookahead);
	} else {
		write_token(out, g, item.lookahead);
	}
}

void write_item(FILE *out, const Grammar *g, const Item item) {
	write_item_helper(out, g, item, 0);
}

void write_item_escaped(FILE *out, const Grammar *g, const Item item) {
	write_item_helper(out, g, item, 1);
}
