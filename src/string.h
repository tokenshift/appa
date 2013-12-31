#ifndef STRING_H
#define STRING_H

#include <stdlib.h>

// TODO: Support different character encodings (ASCII, UTF-8, wchar_t).
typedef char t_char;

typedef struct {
	size_t len;
	const t_char *val;
} String;

String str(const t_char *val);
int str_cmp(String a, String b);
const t_char *str_val(const String str);

#endif
