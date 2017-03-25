#include <stdio.h>
#include "scanner.h"

extern int yylex();
extern int yylineno;
extern char* yytext;

char *names[] = {NULL,
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
        "assignment"};

int lexer(void) {
    int ntoken, vtoken;

    ntoken = yylex();

    while(ntoken) {
        printf("%d - %s\n", ntoken, names[ntoken]);
        ntoken = yylex();
    }
    return 0;
}
