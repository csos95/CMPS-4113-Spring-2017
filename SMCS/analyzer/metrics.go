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
		if token.Type == "NEWLINE" {
			if nl == false {
				nl = true
				lines++
			}
		} else if token.Type == "LINE_COMMENT" || token.Type == "BLOCK_COMMENT" {
			nl = false
			lines--
		} else {
			nl = false
		}
	}
	if lines == 1 {
		return Result{Metric: "Lines of Code", Body: template.HTML(fmt.Sprintf("There is %d line of code.", lines))}, nil
	}
	return Result{Metric: "Lines of Code", Body: template.HTML(fmt.Sprintf("There are %d lines of code.", lines))}, nil
}

func LinesOfDocumentation(tokens []Token) (Result, error) {
	lines := 0
	for _, token := range tokens {
		if token.Type == "LINE_COMMENT" {
			lines++
		} else if token.Type == "BLOCK_COMMENT" {
			lines += strings.Count(token.Value, "\n") + 1
		}
	}
	if lines == 1 {
		return Result{Metric: "Lines of Documentation", Body: template.HTML(fmt.Sprintf("There is %d line of documentation.", lines))}, nil
	}
	return Result{Metric: "Lines of Documentation", Body: template.HTML(fmt.Sprintf("There are %d lines of documentation.", lines))}, nil
}

func BlankLines(tokens []Token) (Result, error) {
	lines := 0
	nl := false
	for _, token := range tokens {
		if token.Type == "NEWLINE" {
			if nl == false {
				nl = true
			} else {
				lines++
			}
		} else {
			nl = false
		}
	}
	if lines == 1 {
		return Result{Metric: "Blank Lines", Body: template.HTML(fmt.Sprintf("There is %d blank line.", lines))}, nil
	}
	return Result{Metric: "Blank Lines", Body: template.HTML(fmt.Sprintf("There are %d blank lines.", lines))}, nil
}

func TotalLines(tokens []Token) (Result, error) {
	lines := 1

	for _, token := range tokens {
		if token.Type == "NEWLINE" {
			lines++
		}
	}
	if lines == 1 {
		return Result{Metric: "Total Lines", Body: template.HTML(fmt.Sprintf("There is %d line total.", lines))}, nil
	}
	return Result{Metric: "Total Lines", Body: template.HTML(fmt.Sprintf("There are %d lines total.", lines))}, nil
}

func NumberOfFunctions(tokens []Token) (Result, error) {
	funcs := 0
	for _, token := range tokens {
		if token.Type == "FUNCTION" {
			funcs++
		}
	}
	if funcs == 1 {
		return Result{Metric: "Number of Functions", Body: template.HTML(fmt.Sprintf("There is %d function.", funcs))}, nil
	}
	return Result{Metric: "Number of Functions", Body: template.HTML(fmt.Sprintf("There are %d functions.", funcs))}, nil
}
