#include <assert.h>
#include <stdarg.h>
#include <stdio.h>

#include "appa.h"
#include "grammar.h"
#include "production.h"
#include "token.h"

Grammar *appa_create_grammar() {
	Grammar *g = malloc(sizeof(Grammar));
	g->tokens = create_vector(sizeof(tkn_info), 32);
	g->productions = create_vector(sizeof(Production), 32);
	return g;
}

void appa_add_rule(Grammar *g, NonTerminal nt, int len, ...) {
	assert(len > 0);
	assert(nt < vec_len(g->tokens));
	assert(((tkn_info *) vec_at(g->tokens, nt))->type == TKN_NONTERM);

	Production *rule = vec_push(g->productions);
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
	tkn_info *term = vec_push(g->tokens);
	term->type = TKN_LITERAL;
	term->value = value;

	return vec_len(g->tokens) - 1;
}

NonTerminal appa_nonterminal(Grammar *g, String name) {
	int i;
	for (i = 0; i < vec_len(g->tokens); ++i) {
		tkn_info *tkn = vec_at(g->tokens, i);
		if (tkn->type == TKN_NONTERM &&
			str_cmp(name, tkn->name) == 0) {
			return i;
		}
	}

	tkn_info *nt = vec_push(g->tokens);
	nt->type = TKN_NONTERM;
	nt->name = name;
	nt->productions = create_vector(sizeof(Vector *), 4);

	return vec_len(g->tokens) - 1;
}

void write_rule(const Grammar *g, Production *rule, FILE *out) {
	tkn_info *nt = token_at(g, rule->head);
	fprintf(out, "%s =>", str_val(nt->name));

	int i;
	for (i = 0; i < rule->len; ++i) {
		fprintf(out, " ");
		write_token(out, g, rule->tail[i]);
	}
	fprintf(out, "\n");
}

void appa_write_grammar(const Grammar *g, FILE *out) {
	int i;
	for (i = 0; i < vec_len(g->productions); ++i) {
		Production *rule = vec_at(g->productions, i);
		write_rule(g, rule, out);
	}
}

tkn_info *token_at(const Grammar *g, int index) {
	return ((tkn_info *) vec_at(g->tokens, index));
}
