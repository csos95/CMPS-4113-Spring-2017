package analyzer

import (
	"html/template"
	"fmt"
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
	return Result{Metric: "Lines of Documentation", Body: template.HTML("Not implemented")}, nil
}