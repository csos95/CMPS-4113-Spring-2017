package analyzer

import (
	"testing"
	"fmt"
)

func Test_JAVA1(t *testing.T) {
	program := `//Lines of Code: 10
			//Lines of Documentation: 5
			//Blank lines: 2
			//Total Lines: 17
			//Number of functions: 1

			import java.util.*;
			import java.io.BufferedWriter;
			import java.io.File;

			public class JAVA1
			{
			    public static void main(String[] args)
			    {
				System.out.println("Hello World");
			    }
			}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 10
		case "Lines of Documentation":
			expected = 5
		case "Blank Lines":
			expected = 2
		case "Total Lines":
			expected = 17
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_JAVA2(t *testing.T) {
	program := `//Lines of Code: 14
			//Lines of Documentation: 6
			//Blank lines: 3
			//Total Lines: 23
			//Number of functions: 2

			import java.util.*;
			import java.io.BufferedWriter;
			import java.io.File;

			public class JAVA1
			{
			    public static void main(String[] args)
			    {
				printHello();
			    }

			    //print hello function
			    public static void printHello()
			    {
				System.out.println("Hello World");
			    }
			}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 14
		case "Lines of Documentation":
			expected = 6
		case "Blank Lines":
			expected = 3
		case "Total Lines":
			expected = 23
		case "Number of Functions":
			expected = 2
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_JAVA3(t *testing.T) {
	program := `//Lines of Code: 32
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
			}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 32
		case "Lines of Documentation":
			expected = 11
		case "Blank Lines":
			expected = 8
		case "Total Lines":
			expected = 51
		case "Number of Functions":
			expected = 5
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_JAVA_Hello(t *testing.T) {
	program := `public class HelloWorld {
		public static void main(String[] args) {
			System.out.println("Hello, World!");
		}
	}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 5
		case "Lines of Documentation":
			expected = 0
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 5
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_JAVA_Add(t *testing.T) {
	program := `public class HelloWorld {
		public static void main(String[] args) {
			int x = 1;
			int y = 2;
			int z = x + y;
			System.out.println(z);
		}
	}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 8
		case "Lines of Documentation":
			expected = 0
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 8
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_JAVA_Document(t *testing.T) {
	program := `/*This is a super awesome program that does cool stuff.
	By: [user]*/
	public class HelloWorld {
		public static void main(String[] args) {
			int x = 1;
			int y = 2;
			int z = x + y;
			//useless comment
			System.out.println(z);
		}
	}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 8
		case "Lines of Documentation":
			expected = 3
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 11
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}
