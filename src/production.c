#include "production.h"
#include "set.h"

int hash_prod(const void *prod) {
	Production *p = (Production *) prod;
	int h = hash(hash_init, p->head);
	h = hash(h, p->len);
	
	int i;
	for (i = 0; i < p->len; ++i) {
		h = hash(h, p->tail[i]);
	}

	return h;
}
