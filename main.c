#include <stdio.h>

#include "src/appa.h"

void test_grammar_1() {
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

	appa_write_dot_grammar(g, stdout, exp);
	appa_delete_grammar(g);
}

void test_grammar_2() {
	Grammar *g = appa_create_grammar();

	Terminal op1 = appa_regex(g, str("+|-"));
	Terminal op2 = appa_regex(g, str("\\*|\\/"));
	Terminal num = appa_regex(g, str("\\d+"));
	Terminal lparen = appa_literal(g, str("("));
	Terminal rparen = appa_literal(g, str(")"));

	NonTerminal exp = appa_nonterminal(g, str("E"));
	NonTerminal term = appa_nonterminal(g, str("T"));
	NonTerminal val = appa_nonterminal(g, str("V"));

	appa_add_rule(g, exp, 1, term);
	appa_add_rule(g, exp, 3, exp, op1, term);

	appa_add_rule(g, term, 1, val);
	appa_add_rule(g, term, 3, term, op2, val);

	appa_add_rule(g, val, 1, num);
	appa_add_rule(g, val, 3, lparen, exp, rparen);

	appa_write_dot_grammar(g, stdout, exp);
	appa_delete_grammar(g);
}

int main(int argc, char **argv) {
	test_grammar_2();

	return 0;
}
