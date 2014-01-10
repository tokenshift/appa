#include <assert.h>
#include <stdio.h>

#include "grammar.h"
#include "item.h"
#include "token.h"

void write_token(FILE *out, const Grammar *g, Token t) {
	token *tkn = token_info(g, t);
	switch (tkn->type) {
		case TKN_START:
			fprintf(out, "S");
			break;
		case TKN_EOF:
			fprintf(out, "$");
			break;
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
