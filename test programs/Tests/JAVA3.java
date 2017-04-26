//Lines of Code: 32
//Lines of Documentation:11
//Blank lines: 8
//Total Lines: 51
//Number of functions: 5

import java.util.*;
import java.io.BufferedWriter;
import java.io.File;

public class JAVA1
{
    public static void main(String[] args)
    {
        //Variable declarations
        int a = 1, b = 5;
        double c = 0.323, d, 43.532;
        float x = 0.00031, y = 421.421;

        printHello();

        //Function calls
        System.out.println("%d + %d = %d", a, b, add(a, b));
        System.out.println("%f / %f = %f", d, c, divide(d, c));
        System.out.println("%f * %f = %f", x, y, multiply(x, y));
    }

    //print hello function
    public static void printHello()
    {
        System.out.println("Hello World");
    }

    //Add function
    public static int add(int a, int b)
    {
        return a + b;
    }

    //Divide function
    public static double divide(double a, double b)
    {
        return a / b;
    }

    //Multiply function
    public static float multiply(float a, float b)
    {
        return a * b;
    }
}