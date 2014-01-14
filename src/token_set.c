#include "token_set.h"
#include "vector.h"

struct TokenSet {
	Vector *tokens;
};

TokenSet *token_set_new() {
	TokenSet *set = calloc(1, sizeof(TokenSet));
	set->tokens = vec_new(sizeof(Token), 4);
	return set;
}

void token_set_delete(TokenSet *set) {
	vec_delete(set->tokens);
	free(set);
}

void token_set_add(TokenSet *set, Token t) {
	int i;
	for (i = 0; i < vec_len(set->tokens); ++i) {
		if (*(Token *)vec_at(set->tokens, i) == t) return;
	}

	*(Token *)vec_push(set->tokens) = t;
}

Token token_set_at(const TokenSet *set, int index) {
	return *(Token *)vec_at(set->tokens, index);
}

int token_set_len(const TokenSet *set) {
	return vec_len(set->tokens);
}
