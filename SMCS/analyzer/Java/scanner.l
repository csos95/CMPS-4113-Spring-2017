%{
#include "scanner.h"
%}

%%

"import java."+*                    return IMPORT;
"public"                            return PUBLIC;
"private"                           return PRIVATE;
"protected"                         return PROTECTED;
"class"                             return CLASS;

"int"                               return tINT;
"float"|"double"                    return tFLOAT;
"char"                              return tCHAR;
"string"                            return tSTRING;
"bool"                              return tBOOL;
"static"                            return tSTATIC;
"void"                              return tVOID;
"final"                             return tFINAL;
"const"                             return tCONST;

[0-9]+                              return INT;
[0-9]+.[0-9]+                       return FLOAT;
"'"\w|\W"'"                         return CHAR;
\".*\"                              return STRING;
"true"|"false"                      return BOOL;
"null"                              return NULL;

";"                                 return SEMICOLON;
":"                                 return COLON;
"{"                                 return LEFT_BRACE;
"}"                                 return RIGHT_BRACE;
"return"                            return RETURN;
(static)?[a-zA-Z]+" "[a-zA-Z]+" "?\(.*\)    return FUNCTION;
[a-zA-Z]+\(.*\)                     return FUNCTION_CALL;
\/\/.*                              return LINE_COMMENT;
\/\*(?:[^\*]|\*[^\/])*\*\/          return BLOCK_COMMENT;
[_a-zA-Z][_a-zA-Z0-9]*              return VAR_NAME;

//relation operators
"=="                                return EQ;
"!="                                return NE;
">"                                 return GT;
"<"                                 return LT;
"<="                                return LE;
">="                                return GE;

"*"                                 return MUL;
"+"                                 return ADD;
"-"                                 return SUB;
"/"                                 return DIV;
"%"                                 return MOD;
"++"                                return INC;
"--"                                return DEC;

"="                                 return ASSIGNMENT;
"!"                                 return NOT;
"||"                                return OR;
"&&"                                return AND;
"+="                                return ADDE;
"-="                                return SUBE;
"*="                                return MULE;
"/="                                return DIVE;
"%="                                return MODE;
"&"                                 return BAND;
"|"                                 return BOR;
"^"                                 return BXOR;
"~"                                 return BNOT;
"<<"                                return BSHL;
">>"                                return BSHR;




[ \t\n]                             ;
.                                   printf("wat.");
%%

int yywrap(void) {
    return 1;
}
