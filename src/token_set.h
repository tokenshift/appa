#ifndef TOKEN_SET_H
#define TOKEN_SET_H

#include "token.h"

typedef struct TokenSet TokenSet;

// Creates a new token set.
TokenSet *token_set_new();

// Deletes a token set.
void token_set_delete(TokenSet *set);

// Adds a token to the set.
void token_set_add(TokenSet *set, Token t);

// Gets the token at the specified index.
Token token_set_at(const TokenSet *set, int index);

// Gets the number of tokens in the set.
int token_set_len(const TokenSet *set);

#endif
