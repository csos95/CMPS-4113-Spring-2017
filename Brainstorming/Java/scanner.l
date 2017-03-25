%{
#include "scanner.h"
%}

%%

"{"                         return LEFT_BRACE;
"}"                         return RIGHT_BRACE;

[0-9]+                      return INT;
[0-9]+.[0-9]+               return FLOAT;
"'"."'"                     return CHAR;
\".*\"                      return STRING;
"#include"                  return INCLUDE;
"<"[a-zA-Z]+(.h)?">"        return HEADER;
"return"                    return RETURN;
;                           return SEMICOLON;
[a-zA-Z]+" "[a-zA-Z]+\(.*\)                  return FUNCTION;
[a-zA-Z]+\(.*\)             return FUNCTION_CALL;

"int"                       return tINT;
"float"                     return tFLOAT;
"char"                      return tCHAR;
"string"                    return tSTRING;

"="                         return ASSIGNMENT;
"*"                         return MUL;
"+"                         return ADD;
"-"                         return SUB;

[_a-zA-Z][_a-zA-Z0-9]*      return VAR_NAME;

[ \t\n]                             ;
.                                   printf("wat.");
%%

int yywrap(void) {
    return 1;
}
