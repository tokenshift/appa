#include "production.h"
#include "set.h"
#include "token.h"

int comp_prod(const void *a, const void *b) {
	const production *p1, *p2;
	p1 = a; p2 = b;

	if (p1->head != p2->head) {
		return p1->head - p2->head;
	}

	if (p1->len != p2->len) {
		return p1->len - p2->len;
	}

	int i;
	for (i = 0; i < p1->len; ++i) {
		int diff = p1->tail[i] - p2->tail[i];
		if (diff != 0) {
			return diff;
		}
	}

	return 0;
}

int hash_prod(const void *prod) {
	const production *p = prod;
	int h = hash(hash_init, p->head);
	h = hash(h, p->len);
	
	int i;
	for (i = 0; i < p->len; ++i) {
		h = hash(h, p->tail[i]);
	}

	return h;
}

void write_production(FILE *out, const Grammar *g, production *rule) {
	token *nt = token_info(g, rule->head);
	fprintf(out, "%s =>", str_val(nt->name));

	int i;
	for (i = 0; i < rule->len; ++i) {
		fprintf(out, " ");
		write_token(out, g, rule->tail[i]);
	}
	fprintf(out, "\n");
}
