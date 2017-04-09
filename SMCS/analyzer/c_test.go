package analyzer

import (
	"testing"
	"fmt"
)

func Test_C_Hello(t *testing.T) {
	program := `#include <stdio.h>
	int main() {
		printf("Hello, World!\n");
		return 0;
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

func Test_C_Add(t *testing.T) {
	program := `#include <stdio.h>
	int main() {
		int x = 1;
		int y = 2;
		int z = x + y;
		printf("%d\n", z);
		return 0;
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

func Test_C_Document(t *testing.T) {
	program := `/*This is a super awesome program that does cool stuff.
	By: [user]*/
	#include <stdio.h>
	int main() {
		int x = 1;
		int y = 2;
		int z = x + y;
		printf("%d\n", z);
		//useless comment
		return 0;
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