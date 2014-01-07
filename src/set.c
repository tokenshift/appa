#include <assert.h>

#include "set.h"

typedef struct set_link set_link;

struct set_link {
	void *item;
	set_link *next;
};

struct Set {
	size_t size;
	set_hash_fun hash;
	set_comp_fun comp;
	set_link *sets;
};

Set *create_set(size_t size, set_hash_fun hash, set_comp_fun comp) {
	assert(size > 0);

	Set *set = malloc(sizeof(Set) + size*sizeof(set_link));
	set->size = size;
	set->hash = hash;
	set->comp = comp;
	set->sets = (set_link *) (set + sizeof(Set));

	int i;
	for (i = 0; i < size; ++i) {
		set->sets[i].item = 0;
		set->sets[i].next = 0;
	}

	return set;
}

void delete_set(Set *set) {
	free(set);
}

void *find_or_put(Set *set, const void *find, void *put) {
	int hash = set->hash(find);
	hash = hash % set->size;

	set_link *link;
	set_link *last_link = 0;
	for (link = &set->sets[hash]; link != 0 && link->item != 0; link = link->next) {
		if (set->comp(link->item, find) == 0) {
			return link->item;
		}

		last_link = link;
	}

	if (put != 0) {
		if (last_link == 0) {
			link = &set->sets[hash];
		} else {
			link = last_link->next = malloc(sizeof(set_link));
		}
		
		link->item = put;
		link->next = 0;

		return put;
	}

	return 0;
}

void *set_find(Set *set, const void *item) {
	return find_or_put(set, item, 0);
}

int set_has(Set *set, const void *item) {
	return set_find(set, item) != 0;
}

void *set_pop(Set *set) {
	int i;
	for (i = 0; i < set->size; ++i) {
		set_link *link;
		set_link *last_link = 0;
		for (link = &set->sets[i]; link != 0; link = link->next) {
			if (link->item != 0 && link->next == 0) {
				void *res = link->item;

				if (last_link != 0) {
					last_link->next = 0;
					free(link);
				} else {
					link->item = 0;
				}

				return res;
			}

			last_link = link;
		}
	}

	return 0;
}

void *set_put(Set *set, void *item) {
	return find_or_put(set, item, item);
}

const int hash_init = 2166136261;
int hash(int hash, int val) {
	return (hash * 16777619) ^ val;
}
