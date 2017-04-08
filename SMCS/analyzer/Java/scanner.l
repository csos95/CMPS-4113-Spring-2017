%{
#include "scanner.h"
%}

%%

"import java".*";"                  return IMPORT;
(public|private|protected)" class "[a-zA-Z]+       return CLASS;
(public|private|protected)" interface "[a-zA-Z]+    return INTERFACE;
"new"                               return NEW;
","                                 return COMMA;

"if"[ \t]*\(.*\)                    return IF;
"else"                              return ELSE;
"while"[ \t]*\(.*\)                 return WHILE;
"for"[ \t]*\(.*\)                   return FOR;
"switch"[ \t]*\(.*\)                return SWITCH;
"do"                                return DO;
"try"                               return TRY;
"catch"[ \t]*\(.*\)                 return CATCH;
"case"[ \t]+.+":"                   return CASE;
"break"                             return BREAK;
"goto"                              return GOTO;
"continue"                          return CONTINUE;
"throw"[ \t]+.+";"|"throws"([ \t]+[_a-zA-Z][_a-zA-Z0-9]*",")*[ \t]+[_a-zA-Z][_a-zA-Z0-9]*       return THROW;
"finally"                           return FINALLY;

"int"                               return tINT;
"float"|"double"                    return tFLOAT;
"char"                              return tCHAR;
"String"                            return tSTRING;
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
\/\*([^*]|[\r\n]|(\*+([^*\/]|[\r\n])))*\*+\/   return BLOCK_COMMENT;
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
