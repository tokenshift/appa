#ifndef APPA_H
#define APPA_H

#include <stdio.h>

#include "string.h"

/* Appa Public API */

typedef int Token;
typedef int Terminal;
typedef int NonTerminal;

typedef struct Grammar Grammar;
typedef struct Parser Parser;

// Creates a new grammar.
Grammar *appa_create_grammar();

// Deletes a grammar, releasing allocated memory.
void appa_delete_grammar(Grammar *g);

// Constructs the parse table for the grammar.
Parser *appa_compile(const Grammar *g, NonTerminal start);

// Adds a literal to the grammar.
Terminal appa_literal(Grammar *grammar, String value);

// Adds a non-terminal symbol to the grammar.
NonTerminal appa_nonterminal(Grammar *grammar, String name);

// Adds a production to the specified non-terminal symbol.
void appa_add_rule(Grammar *grammar, NonTerminal head, int len, ...);

// Writes all of the grammar's production rules to the specified output stream.
void appa_write_grammar(const Grammar *g, FILE *out);

// Writes the LALR automaton in Graphviz .dot format.
void appa_write_dot_grammar(const Grammar *g, FILE *out, NonTerminal start);

#endif
