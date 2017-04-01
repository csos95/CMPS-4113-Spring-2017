%{
#include "scanner.h"
%}

%%

"import java."+*                    return IMPORT;
(public|private|protected)" class "[a-zA-Z]+       return CLASS;
(public|private|protected)" interface "[a-zA-Z]+    return INTERFACE;

"int"                               return tINT;
"float"|"double"                    return tFLOAT;
"char"                              return tCHAR;
"string"                            return tSTRING;
"bool"                              return tBOOL;
"static"                            return tSTATIC;
"void"                              return tVOID;
"const"|"final"                     return tCONST;

[0-9]+                              return INT;
[0-9]+.[0-9]+                       return FLOAT;
"'"\w|\W"'"                         return CHAR;
\".*\"                              return STRING;
"true"|"false"                      return BOOL;

";"                                 return SEMICOLON;
":"                                 return COLON;
"{"                                 return LEFT_BRACE;
"}"                                 return RIGHT_BRACE;
"return"                            return RETURN;
(public|private|protected)[ \t]+(static)?[ \t]+[a-zA-Z]+[ \t]+[a-zA-Z]+[ \t]*\(.*\)    return FUNCTION;
[a-zA-Z\.]+\(.*\)                     return FUNCTION_CALL;
\/\/.*                              return LINE_COMMENT;
\/*([^*]|[\r\n]|(\*+([^*\/]|[\r\n])))*\*+\/   return BLOCK_COMMENT;
[_a-zA-Z][_a-zA-Z0-9]*              return VAR_NAME;

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



"\r"?"\n"			return NEWLINE;
[ \t]                             ;
.                                   printf("wat.");
%%

int yywrap(void) {
    return 1;
}
