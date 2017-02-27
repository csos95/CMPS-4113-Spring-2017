package c

/*
#include <stdio.h>
#include "scanner.h"

typedef struct yy_buffer_state *YY_BUFFER_STATE;
extern YY_BUFFER_STATE yy_scan_string(char*);
extern void yy_delete_buffer(YY_BUFFER_STATE);

extern int yylex(void);
*/
import "C"

var names = []string{"NULL",
	"int",
	"char",
	"float",
	"string",
	"bool",
	"",
	"",
	"",
	"",
	"",
	"int type",
	"char type",
	"float type",
	"string type",
	"bool type",
	"",
	"",
	"",
	"",
	"",
	"semicolon",
	"colon",
	"left brace",
	"right brace",
	"double quote",
	"single quote",
	"newline",
	"",
	"",
	"",
	"function",
	"variable name",
	"include",
	"header",
	"identifier",
	"function call",
	"return",
	"line comment",
	"block comment",
	"",
	"EQ",
	"NE",
	"GT",
	"LT",
	"GE",
	"LE",
	"",
	"",
	"",
	"",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"INC",
	"DEC",
	"",
	"",
	"",
	"assignment"}

var state *C.struct_yy_buffer_state

func Parse(source string) {
	state = C.yy_scan_string(C.CString(source))
}

func NextToken() (string, string) {
	return names[C.int(C.yylex())], ""
}

func Close() {
	C.yy_delete_buffer(state)
}
