package analyzer

import (
	"testing"
	"fmt"
)

func Test_CPP_Hello(t *testing.T) {
	program := `#include <iostream>
	using namespace std;
	int main() {
		cout << "Hello, World!" << endl;
		return 0;
	}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 6
		case "Lines of Documentation":
			expected = 0
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 6
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d, got %d", expected, result.Value))
		}
	}
}

func Test_CPP_Add(t *testing.T) {
	program := `#include <iostream>
	using namespace std;
	int main() {
		int x = 1;
		int y = 2;
		int z = x + y;
		cout << z << endl;
		return 0;
	}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 9
		case "Lines of Documentation":
			expected = 0
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 9
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d, got %d", expected, result.Value))
		}
	}
}

func Test_CPP_Document(t *testing.T) {
	program := `/*This is a super awesome program that does cool stuff.
	By: [user]*/
	#include <iostream>
	using namespace std;
	int main() {
		int x = 1;
		int y = 2;
		int z = x + y;
		cout << z << endl;
		//useless comment
		return 0;
	}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 9
		case "Lines of Documentation":
			expected = 3
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 12
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d, got %d", expected, result.Value))
		}
	}
}