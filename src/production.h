#ifndef PRODUCTION_H
#define PRODUCTION_H

#include "appa.h"

typedef struct {
	NonTerminal head;
	int len;
	Token *tail;
} production;

int comp_prod(const void *a, const void *b);
int hash_prod(const void *prod);
void write_production(FILE *out, const Grammar *g, production *rule);

#endif
