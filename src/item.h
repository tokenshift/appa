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

// Checks whether two item cores (ignoring lookaheads) are identical.
int item_core_eq(const Item a, const Item b);

// Checks whether two items are identical.
int item_eq(const Item a, const Item b);

void write_item(FILE *out, const Grammar *g, const Item item);
void write_item_escaped(FILE *out, const Grammar *g, const Item item);

#endif
