package main

/*
#include <stdio.h>
#include "myscanner.h"

extern int lexer(void);
*/
import "C"

func main() {
    C.lexer();
}
