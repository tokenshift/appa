#ifndef ITEM_H
#define ITEM_H

#include <stdio.h>

#include "appa.h"
#include "production.h"

typedef struct {
	production *rule;
	int pos;
	Token lookahead;
} Item;

void write_item(FILE *out, const Grammar *g, const Item item);

int hash_item(const void *item);
int comp_item(const void *a, const void *b);

#endif
