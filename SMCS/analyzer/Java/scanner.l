%{
#include "scanner.h"
%}

%%



[ \t\n]                             ;
.                                   printf("wat.");
%%

int yywrap(void) {
    return 1;
}
