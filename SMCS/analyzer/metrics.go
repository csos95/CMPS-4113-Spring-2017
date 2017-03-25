package analyzer

import (
	"fmt"
	"html/template"
	"strings"
)

type Metric func([]Token) (Result, error)

func LinesOfCode(tokens []Token) (Result, error) {
	lines := 1
	nl := false
	for _, token := range tokens {
		if token.Type == "newline" {
			if nl == false {
				nl = true
				lines++
			}
		} else if token.Type == "line comment" || token.Type == "block comment" {
			nl = false
			lines--
		} else {
			nl = false
		}
	}
	return Result{Metric: "Lines of Code", Body: template.HTML(fmt.Sprintf("There are %d lines of code.", lines))}, nil
}

func LinesOfDocumentation(tokens []Token) (Result, error) {
	lines := 0
	for _, token := range tokens {
		if token.Type == "line comment" {
			lines++
		} else if token.Type == "block comment" {
			lines += strings.Count(token.Value, "\n") + 1
		}
	}
	return Result{Metric: "Lines of Documentation", Body: template.HTML(fmt.Sprintf("There are %d lines of documentation.", lines))}, nil
}

func BlankLines (tokens []Token) (Result, error) {
	lines := 0
	nl := false
	for _, token := range tokens {
		if token.Type == "newline" {
			if nl == false {
				nl = true
			} else {
				lines++
			}
		} else {
			nl = false
		}
	}
	return Result{Metric: "Blank Lines", Body: template.HTML(fmt.Sprintf("There are %d blank lines.", lines))}, nil
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
