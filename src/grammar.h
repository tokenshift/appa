#ifndef GRAMMAR_H
#define GRAMMAR_H

#include "appa.h"
#include "string.h"
#include "vector.h"

#define TKN_NONTERM	0
#define TKN_LITERAL	1
#define TKN_REGEX	2

typedef struct {
	int type;
	union {
		String name;
		String value;
		Vector *productions;
	};
} tkn_info;

tkn_info *token_at(const Grammar *g, int index);

#endif
