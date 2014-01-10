#include <assert.h>
#include <stdarg.h>
#include <stdio.h>

#include "appa.h"
#include "grammar.h"
#include "production.h"
#include "token.h"

Grammar *appa_create_grammar() {
	Grammar *g = malloc(sizeof(Grammar));
	g->tokens = create_vector(sizeof(token), 8);
	g->productions = create_vector(sizeof(production), 8);
	return g;
}

void appa_add_rule(Grammar *g, NonTerminal nt, int len, ...) {
	assert(len > 0);
	assert(nt < vec_len(g->tokens));
	assert(((token *) vec_at(g->tokens, nt))->type == TKN_NONTERM);

	production *rule = vec_push(g->productions);
	rule->head = nt;
	rule->len = len;
	rule->tail = calloc(len, sizeof(Token));

	va_list args;
	va_start(args, len);

	int i;
	for (i = 0; i < len; ++i) {
		rule->tail[i] = va_arg(args, Token);
	}

	va_end(args);
}

Terminal appa_literal(Grammar *g, String value) {
	token *term = vec_push(g->tokens);
	term->type = TKN_LITERAL;
	term->value = value;

	return vec_len(g->tokens) - 1;
}

NonTerminal appa_nonterminal(Grammar *g, String name) {
	int i;
	for (i = 0; i < vec_len(g->tokens); ++i) {
		token *tkn = vec_at(g->tokens, i);
		if (tkn->type == TKN_NONTERM &&
			str_cmp(name, tkn->name) == 0) {
			return i;
		}
	}

	token *nt = vec_push(g->tokens);
	nt->type = TKN_NONTERM;
	nt->name = name;

	return vec_len(g->tokens) - 1;
}

void appa_write_grammar(const Grammar *g, FILE *out) {
	int i;
	for (i = 0; i < vec_len(g->productions); ++i) {
		production *rule = vec_at(g->productions, i);
		write_production(out, g, rule);
	}
}

token *token_info(const Grammar *g, Token t) {
	assert(t >= 0);
	assert(t < vec_len(g->tokens));

	return vec_at(g->tokens, t);
}
