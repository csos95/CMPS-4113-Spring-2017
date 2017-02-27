package analyzer

import (
	"fmt"
	"html/template"
)

type Metric func([]Token) (Result, error)

func LinesOfCode(tokens []Token) (Result, error) {
	lines := 1
	for _, token := range tokens {
		if token.Type == "newline" {
			lines++
		}
	}
	return Result{Metric: "Lines of Code", Body: template.HTML(fmt.Sprintf("There are %d lines of code.", lines))}, nil
}

func LinesOfDocumentation(tokens []Token) (Result, error) {
	lines := 0
	for _, token := range tokens {
		if token.Type == "line comment" || token.Type == "block comment" {
			lines++
		}
	}
	return Result{Metric: "Lines of Documentation", Body: template.HTML(fmt.Sprintf("There are %d instances of documentation (lines coming soon).", lines))}, nil
}

func NumberOfFunctions(tokens []Token) (Result, error) {
	funcs := 0
	for _, token := range tokens {
		if token.Type == "function" {
			funcs++
		}
	}
	return Result{Metric: "Number of Functions", Body: template.HTML(fmt.Sprintf("There are %d functions.", funcs))}, nil
}
