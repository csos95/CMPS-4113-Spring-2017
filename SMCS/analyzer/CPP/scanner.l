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
\/\/.*                      return LINE_COMMENT;
\/\*(?:[^\*]|\*[^\/])*\*\/  return BLOCK_COMMENT;
"#include"                  return INCLUDE;
"<"[a-zA-Z]+(.h)?">"        return HEADER;
"return"                    return RETURN;
;                           return SEMICOLON;
[a-zA-Z]+" "[a-zA-Z]+" "?\(.*\) return FUNCTION;
[a-zA-Z]+\(.*\)             return FUNCTION_CALL;

"int"                       return tINT;
"float"                     return tFLOAT;
"char"                      return tCHAR;
"string"                    return tSTRING;

"="                         return ASSIGNMENT;
"*"                         return MUL;
"+"                         return ADD;
"-"                         return SUB;
"!"                         return NOT;
"||"                        return OR;
"&&"                        return AND;
"+="                        return ADDE;
"-="                        return SUBE;
"*="                        return MULE;
"/="                        return DIVE;
"%="                        return MODE;
"&"                         return BADD;
"|"                         return BOR;
"^"                         return BXOR;
"~"                         return BNOT;
"<<"                        return BSHL;
">>"                        return BSHR;
"if"                        return IF;
"else"                      return ELSE;
"While"                     return WHILE;
"for"                       return FOR;
"switch"                    return SWITCH;
"case"                      return CASE;
"break"                     return BREAK;
"continue"                  return CONTINUE;
"finally"                   return FINALLY;

[_a-zA-Z][_a-zA-Z0-9]*      return VAR_NAME;

"\r"?"\n"                        return NEWLINE;

[ \t]                             ;
.                                   printf("wat.");
%%

int yywrap(void) {
    return 1;
}
