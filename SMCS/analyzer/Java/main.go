package main
import (
    "fmt"
    "io/ioutil"
)

/*
#include <stdio.h>
#include "scanner.h"

typedef struct yy_buffer_state *YY_BUFFER_STATE;
extern YY_BUFFER_STATE yy_scan_string(char*);
extern void yy_delete_buffer(YY_BUFFER_STATE);

extern int yylex(void);
*/
import "C"

var names = []string{
        "",
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
        "BSHR"};

func main() {
    file, err := ioutil.ReadFile("input.java")
    if err != nil {
        fmt.Println(err)
    }

    contents := string(file)
    state := C.yy_scan_string(C.CString(contents))

    ntoken := C.int(C.yylex())

    for ntoken != 0 {
        fmt.Printf("%d - %s\n", ntoken, names[ntoken])
        ntoken = C.int(C.yylex())
    }

    C.yy_delete_buffer(state);

}

