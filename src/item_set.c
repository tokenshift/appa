#include <assert.h>

#include "item_set.h"
#include "vector.h"

struct ItemSet {
	Vector *items;
};

ItemSet *item_set_new() {
	ItemSet *set = calloc(1, sizeof(ItemSet));
	set->items = vec_new(sizeof(Item), 4);
	return set;
}

void item_set_delete(ItemSet *set) {
	vec_delete(set->items);
	free(set);
}

void item_set_add(ItemSet *set, Item item) {
	int i;
	for (i = 0; i < vec_len(set->items); ++i) {
		Item *existing = vec_at(set->items, i);
		if (item_eq(item, *existing)) return;
	}

	*(Item *)vec_push(set->items) = item;
}

Item item_set_at(const ItemSet *set, int index) {
	assert(index < vec_len(set->items));

	return *(Item *)vec_at(set->items, index);
}

int item_set_len(const ItemSet *set) {
	return vec_len(set->items);
}
