package analyzer

import (
	"fmt"
	"html/template"
	"strings"
)

type Metric func([]Token) (Result, error)

func LinesOfCode(tokens []Token) (Result, error) {
	lines := 0
	codeLine := false

	for _, token := range tokens {
		if token.Type != "LINE_COMMENT" && token.Type != "BLOCK_COMMENT" && token.Type != "NEWLINE" {
			codeLine = true
		}
		if token.Type == "NEWLINE" {
			if codeLine {
				lines++
			}
			codeLine = false
		}
	}

	if codeLine {
		lines++
	}

	if lines == 1 {
		return Result{Metric: "Lines of Code", Value: lines, Body: template.HTML(fmt.Sprintf("There is %d line of code.", lines))}, nil
	}
	return Result{Metric: "Lines of Code", Value: lines, Body: template.HTML(fmt.Sprintf("There are %d lines of code.", lines))}, nil
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
		return Result{Metric: "Lines of Documentation", Value: lines, Body: template.HTML(fmt.Sprintf("There is %d line of documentation.", lines))}, nil
	}
	return Result{Metric: "Lines of Documentation", Value: lines, Body: template.HTML(fmt.Sprintf("There are %d lines of documentation.", lines))}, nil
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
		return Result{Metric: "Blank Lines", Value: lines, Body: template.HTML(fmt.Sprintf("There is %d blank line.", lines))}, nil
	}
	return Result{Metric: "Blank Lines", Value: lines, Body: template.HTML(fmt.Sprintf("There are %d blank lines.", lines))}, nil
}

func TotalLines(tokens []Token) (Result, error) {
	lines := 0

	if len(tokens) != 0 {
		lines = 1
	}

	for _, token := range tokens {
		if token.Type == "NEWLINE" {
			lines++
		} else if token.Type == "BLOCK_COMMENT" {
			lines += strings.Count(token.Value, "\n")
		}
	}
	if lines == 1 {
		return Result{Metric: "Total Lines", Value: lines, Body: template.HTML(fmt.Sprintf("There is %d line total.", lines))}, nil
	}
	return Result{Metric: "Total Lines", Value: lines, Body: template.HTML(fmt.Sprintf("There are %d lines total.", lines))}, nil
}

func NumberOfFunctions(tokens []Token) (Result, error) {
	funcs := 0
	for _, token := range tokens {
		if token.Type == "FUNCTION" {
			funcs++
		}
	}
	if funcs == 1 {
		return Result{Metric: "Number of Functions", Value: funcs, Body: template.HTML(fmt.Sprintf("There is %d function.", funcs))}, nil
	}
	return Result{Metric: "Number of Functions", Value: funcs, Body: template.HTML(fmt.Sprintf("There are %d functions.", funcs))}, nil
}

func NumberOfClasses(tokens []Token) (Result, error) {
	classes := 0

	for _, token := range tokens {
		if token.Type == "CLASS" {
			classes++
		}
	}

	if classes == 1 {
		return Result{Metric: "Number of Classes", Value: classes, Body: template.HTML(fmt.Sprintf("There is %d class.", classes))}, nil
	}
	return Result{Metric: "Number of Classes", Value: classes, Body: template.HTML(fmt.Sprintf("There are %d classes.", classes))}, nil
}

//<!--<input type="checkbox" name="metric" value="Ratio of LOC to LOD">Ratio of LOC to LOD<br>-->
//
//<!--<input type="checkbox" name="metric" value="Number of Function Parameters">Number of Function Parameters<br>-->
//<!--<input type="checkbox" name="metric" value="Methods per Class">Methods per Class<br>-->
//<!--<input type="checkbox" name="metric" value="Lines per Function">Lines per Function<br>-->
//
//<!--<input type="checkbox" name="metric" value="Cyclomatic Complexity">Cyclomatic Complexity<br>-->
