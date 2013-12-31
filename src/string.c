#include <string.h>

#include "string.h"

String str(const t_char *val) {
	String str;
	str.len = strnlen(val, 256);
	str.val = val;

	return str;
}

const t_char *str_val(const String str) {
	return str.val;
}

int str_cmp(String a, String b) {
	int i;
	for (i = 0; i < a.len && i < b.len; ++i) {
		t_char delta = a.val[i] - b.val[i];
		if (delta != 0) {
			return delta;
		}
	}

	if (i < a.len) {
		return -1;
	}

	if (i < b.len) {
		return 1;
	}

	return 0;
}
