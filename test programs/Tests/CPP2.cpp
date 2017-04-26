//Lines of Code: 20
//Lines of Documentation: 10
//Blank lines: 10
//Total Lines: 40
//Number of functions: 2

#include<iostream>
#include<string>
#include<stdio.h>

using namespace std;

//This function adds 2 numbers using a for loop
// and returns their sum
int sum(int a, int b)
{
    int c = a;

    for(int i = 0; i < b; i++){
        c += 1;
    }

    return c;
}

int main()
{
    int a, b, c;
    
    //Hello world
    cout << "Hello World" << endl;

    //Sum integers a and b
    c = sum(a, b);

    //print result
    cout << "Adding 2 integers: " << a << " + " << b << " = " << c << endl;

    return 0;
}