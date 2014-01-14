#ifndef ITEM_SET_H
#define ITEM_SET_H

#include "item.h"

typedef struct ItemSet ItemSet;

// Creates a new item set.
ItemSet *item_set_new();

// Deletes an item set.
void item_set_delete(ItemSet *set);

// Adds an item to the set.
void item_set_add(ItemSet *set, Item item);

// Gets the item at the specified index.
Item item_set_at(const ItemSet *set, int index);

// Gets the number of items in the set.
int item_set_len(const ItemSet *set);

#endif
