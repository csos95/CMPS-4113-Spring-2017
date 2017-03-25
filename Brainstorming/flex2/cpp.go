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
        "assignment",
        "ADDE",
        "SUBE",
        "MULE",
        "DIVE",
        "MODE",
        "",
        "BADD",
        "BOR",
        "BXOR",
        "BNOT",
        "BSHL",
        "BSHR"};

func main() {
    file, err := ioutil.ReadFile("lex.yy.c")
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

