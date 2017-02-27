#include <stdio.h>

//this is a line comment
int main(void) {
    /*
    * this is a block comment
    */
    int x = 2;
    float y = 3.5;
    float z = x * y;
    int w = x + x;
    int u = w - x;
    printf("z: %d\n", z);
    return 0;
}