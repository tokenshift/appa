#ifndef GRAMMAR_H
#define GRAMMAR_H

#include "appa.h"
#include "string.h"
#include "vector.h"

#define TKN_NONTERM 1
#define TKN_LITERAL 2

typedef struct {
	int type;
	union {
		String name;
		String value;
	};
} token;

struct Grammar {
	Vector *tokens;
	Vector *productions;
};

token *token_info(const Grammar *g, Token t);

#endif
