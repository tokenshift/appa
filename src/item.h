#ifndef ITEM_H
#define ITEM_H

#include <stdio.h>

#include "appa.h"
#include "production.h"

extern const int START_SYMBOL;

typedef struct {
	Production *rule;
	int pos;
} Item;

void write_item(FILE *out, const Grammar *g, const Item *item);

int hash_item(const void *item);
int comp_item(const void *a, const void *b);

// Creates an item from the specified production rule.
Item *prod_item(Production *prod, int pos);

#endif
