//Lines of Code: 39
//Lines of Documentation: 12
//Blank lines: 11
//Total Lines: 62
//Number of functions: 6

#include<iostream>
#include<string>
#include<stdio.h>
#include<math.h>

using namespace std;

//Sum function
int sum(int a, int b)
{
    return a + b;
}

//Multiply function
int multiply(int &a, int b)
{
    return a * b;
}

//print function
void printHello()
{
    cout << "Hello y'all" << endl;
}

//Divide function
double divide(double a, double b)
{
    return a / b;
}

//Cosine function
double cosine(float b)
{
    return cos(b);
} 

//Main function
int main()
{
    int a, b, c;
    double x, y, z;
    
    printHello();

    //Bunch of function calls
    a = 1;
    b = 43;
    c = sum(a, b);
    x = cos(3.142);
    y = cos(3.142*3.142);
    z = divide(x, y);
    b = multiply(a, c);

    return 0;
}