#include <assert.h>
#include <stdio.h>

#include "grammar.h"
#include "item.h"
#include "token.h"

void write_token(FILE *out, const Grammar *g, Token t) {
	if (t == START_SYMBOL) {
		fprintf(out, "S");
		return;
	} else if (t == EOF_SYMBOL) {
		fprintf(out, "$");
		return;
	}

	token *tkn = token_info(g, t);

	switch (tkn->type) {
		case TKN_NONTERM:
			fprintf(out, "<%s>", str_val(tkn->name));
			break;
		case TKN_LITERAL:
			fprintf(out, "\"%s\"", str_val(tkn->value));
			break;
		default:
			assert(0);
			break;
	}
}
