#include <assert.h>
#include <stdio.h>

#include "grammar.h"
#include "item.h"
#include "token.h"

void write_token_value(FILE *out, String value, int escape) {
	int i;
	for (i = 0; i < value.len; ++i) {
		if (escape && value.val[i] == '\\') {
			fprintf(out, "\\\\");
		} else {
			fprintf(out, "%c", value.val[i]);
		}
	}
}

void write_token_helper(FILE *out, const Grammar *g, Token t, int escape) {
	token *tkn = token_info(g, t);
	switch (tkn->type) {
		case TKN_START:
			fprintf(out, "S");
			break;
		case TKN_EOF:
			fprintf(out, "$");
			break;
		case TKN_NONTERM:
			fprintf(out, "<");
			write_token_value(out, tkn->name, escape);
			fprintf(out, ">");
			break;
		case TKN_LITERAL:
			fprintf(out, "'");
			write_token_value(out, tkn->value, escape);
			fprintf(out, "'");
			break;
		case TKN_REGEX:
			fprintf(out, "/");
			write_token_value(out, tkn->pattern, escape);
			fprintf(out, "/");
			break;
		default:
			assert(0);
			break;
	}
}

void write_token(FILE *out, const Grammar *g, Token t) {
	write_token_helper(out, g, t, 0);
}

void write_token_escaped(FILE *out, const Grammar *g, Token t) {
	write_token_helper(out, g, t, 1);
}
