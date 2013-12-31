#include <assert.h>
#include <stdio.h>

#include "grammar.h"
#include "token.h"

void write_token(const Grammar *g, Token t, FILE *out) {
	tkn_info *tkn = token_at(g, t);

	switch (tkn->type) {
		case TKN_NONTERM:
			fprintf(out, "<%s>", str_val(tkn->name));
			break;
		case TKN_LITERAL:
			fprintf(out, "\"%s\"", str_val(tkn->value));
			break;
		case TKN_REGEX:
			fprintf(out, "REGEX");
			break;
		default:
			assert(0);
			break;
	}
}
