package analyzer

import (
	"testing"
	"fmt"
)

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
			t.Error(fmt.Sprintf("Expected %d, got %d", expected, result.Value))
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
			t.Error(fmt.Sprintf("Expected %d, got %d", expected, result.Value))
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
			t.Error(fmt.Sprintf("Expected %d, got %d", expected, result.Value))
		}
	}
}
