#ifndef PRODUCTION_H
#define PRODUCTION_H

#include "appa.h"

typedef struct {
	NonTerminal head;
	int len;
	Token *tail;
} production;

void write_production(FILE *out, const Grammar *g, production *rule);

#endif
