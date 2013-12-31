#include <stdio.h>

#include "appa.h"
#include "string.h"

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

	return 0;
}
