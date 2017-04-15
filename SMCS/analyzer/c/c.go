package c

/*
#include <stdio.h>
#include "scanner.h"

typedef struct yy_buffer_state *YY_BUFFER_STATE;
extern YY_BUFFER_STATE yy_scan_string_c(char*);
extern void yy_delete_buffer_c(YY_BUFFER_STATE);

extern int yylex_c(void);
extern int yylineno;
extern char* yytext_c;
*/
import "C"

var names = []string{
	"NULL",
	"INT",
	"CHAR",
	"FLOAT",
	"STRING",
	"BOOL",
	"",
	"",
	"",
	"",
	"",
	"tINT",
	"tCHAR",
	"tFLOAT",
	"tSTRING",
	"tBOOL",
	"tCONST",
	"tVOID",
	"tSTATIC",
	"",
	"",
	"SEMICOLON",
	"COLON",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"NEWLINE",
	"COMMA",
	"",
	"",
	"",
	"",
	"FUNCTION",
	"VAR_NAME",
	"IMPORT",
	"IDENTIFIER",
	"FUNCTION_CALL",
	"RETURN",
	"LINE_COMMENT",
	"BLOCK_COMMENT",
	"",
	"",
	"EQ",
	"NE",
	"GT",
	"LT",
	"GE",
	"LE",
	"NOT",
	"OR",
	"AND",
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
	"ASSIGNMENT",
	"ADDE",
	"MULE",
	"SUBE",
	"DIVE",
	"MODE",
	"",
	"",
	"",
	"",
	"IF",
	"ELSE",
	"WHILE",
	"FOR",
	"SWITCH",
	"DO",
	"TRY",
	"CATCH",
	"CASE",
	"BREAK",
	"GOTO",
	"CONTINUE",
	"THROW",
	"FINALLY",
	"",
	"",
	"",
	"",
	"",
	"",
	"CLASS",
	"STRUCT",
	"INTERFACE",
	"NEW",
	"",
	"",
	"",
	"",
	"",
	"",
	"BAND",
	"BOR",
	"BXOR",
	"BNOT",
	"BSHL",
	"BSHR"}

var state *C.struct_yy_buffer_state

func Parse(source string) {
	state = C.yy_scan_string_c(C.CString(source))
}

func NextToken() (string, string) {
	return names[C.int(C.yylex_c())], C.GoString(C.yytext_c)
}

func Close() {
	C.yy_delete_buffer_c(state)
}
