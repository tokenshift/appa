#include <stdio.h>

#include "src/appa.h"
#include "src/string.h"
#include "src/set.h"

int hash_int(const void *i) {
	return *(int *)i;
}

int comp_int(const void *a, const void *b) {
	return *(int*)a - *(int*)b;
}

int main(int argc, char **argv) {
	Grammar *g = appa_create_grammar();

	Terminal num = appa_literal(g, str("1"));
	Terminal add = appa_literal(g, str("+"));
	Terminal mul = appa_literal(g, str("*"));

	NonTerminal exp = appa_nonterminal(g, str("E"));
	NonTerminal val = appa_nonterminal(g, str("V"));

	appa_add_rule(g, exp, 1, val);
	appa_add_rule(g, exp, 3, exp, add, val);
	
	appa_add_rule(g, val, 1, num);
	appa_add_rule(g, val, 3, val, mul, num);

	appa_write_grammar(g, stdout);

	printf("\n\n");

	appa_compile(g, exp);

	return 0;
}
