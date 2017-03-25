%{
#include "scanner.h"
%}

%%

\/\/.*                      return LINE_COMMENT;
\/\*(?:[^\*]|\*[^\/])*\*\/  return BLOCK_COMMENT;

\".*\"                      return STRING;

"#include"[ \t]*"<"[a-zA-Z]+(.h)?">"        return IMPORT;
"return"                    return RETURN;
;                           return SEMICOLON;

[a-zA-Z]+" "[a-zA-Z]+" "?\(.*\) return FUNCTION;
[a-zA-Z]+\(.*\)             return FUNCTION_CALL;

[_a-zA-Z][_a-zA-Z0-9]*      return VAR_NAME;

[0-9]+                      return INT;
[0-9]+.[0-9]+               return FLOAT;
"'"."'"                     return CHAR;

"int"                       return tINT;
"float"                     return tFLOAT;
"char"                      return tCHAR;
"string"                    return tSTRING;

"="                         return ASSIGNMENT;
"*"                         return MUL;
"+"                         return ADD;
"-"                         return SUB;

"{"                         return LEFT_BRACE;
"}"                         return RIGHT_BRACE;

"\r"?"\n"                        return NEWLINE;

[ \t]                             ;
.                                   printf("wat.");
%%

int yywrap(void) {
    return 1;
}
