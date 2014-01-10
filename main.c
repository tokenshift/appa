#include <stdio.h>

#include "src/appa.h"
#include "src/string.h"
#include "src/set.h"

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

	/*Set *set = create_set(3, sizeof(int), hash_int, comp_int);

	int values[10];
	int i;
	for (i = 0; i < 10; ++i) {
		values[i] = i + 1;
		printf("set_has(%d) => %d\n", values[i], set_has(set, &values[i]));
	}

	for (i = 0; i < 5; ++i) {
		set_put(set, &values[i]);
	}

	for (i = 0; i < 10; ++i) {
		printf("set_has(%d) => %d\n", values[i], set_has(set, &values[i]));
	}*/

	return 0;
}
