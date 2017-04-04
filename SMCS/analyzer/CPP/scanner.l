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
"true"|"false"              return BOOL;
\/\/.*                      return LINE_COMMENT;
\/*([^*]|[\r\n]|(\*+([^*\/]|[\r\n])))*\*+\/   return BLOCK_COMMENT;
"#include"                  return IMPORT;
"<"[a-zA-Z]+(.h)?">"        return IMPORT;
"return"                    return RETURN;
;                           return SEMICOLON;
":"                         return COLON;
[a-zA-Z]+" "[a-zA-Z]+" "?\(.*\) return FUNCTION;
[a-zA-Z]+\(.*\)             return FUNCTION_CALL;

"int"                       return tINT;
"float"                     return tFLOAT;
"char"                      return tCHAR;
"string"                    return tSTRING;
"bool"						return tBOOL;
"const"						return tCONST;	
"void"						return tVOID;
"static"					return tSTATIC;

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
"&"                         return BAND;
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
"do"						return DO;
"try"						return TRY;
"catch"						return CATCH;
"case"                      return CASE;
"break"                     return BREAK;
"goto"						return GOTO;
"continue"                  return CONTINUE;
"throw"						return THROW;
"finally"                   return FINALLY;
"struct"					return STRUCT;
[_a-zA-Z][_a-zA-Z0-9]*      return VAR_NAME;

"\r"?"\n"                        return NEWLINE;

[ \t]                             ;
.                                   printf("wat.");
%%

int yywrap(void) {
    return 1;
}
