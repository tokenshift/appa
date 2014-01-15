#ifndef TOKEN_H
#define TOKEN_H

#include "grammar.h"

void write_token(FILE *out, const Grammar *g, Token tkn);
void write_token_escaped(FILE *out, const Grammar *g, Token tkn);

#endif
