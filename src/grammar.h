#ifndef GRAMMAR_H
#define GRAMMAR_H

#include "appa.h"
#include "string.h"
#include "vector.h"

#define TKN_START 0
#define TKN_EOF 1
#define TKN_NONTERM 2
#define TKN_LITERAL 4
#define TKN_REGEX 8

typedef struct {
	int type;
	union {
		String name;
		String value;
		String pattern;
	};
} token;

struct Grammar {
	Vector *tokens;
	Vector *productions;
};

token *token_info(const Grammar *g, Token t);

#endif
