#include "production.h"
#include "token.h"

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
