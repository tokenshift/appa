#ifndef PRODUCTION_H
#define PRODUCTION_H

#include "appa.h"

typedef struct {
	NonTerminal head;
	int len;
	Token *tail;
} Production;

int hash_prod(const void *prod);

#endif
