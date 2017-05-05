package analyzer

import (
	"fmt"
	"html/template"
	"log"
	"strings"
)

//Metrics is a function type that runs one metric on a slice of tokens and returns the result
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

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % t
		a = t
	}
	return a
}

func RatioOfLOCToLOD(tokens []Token) (Result, error) {
	locResult, _ := LinesOfCode(tokens)
	loc := locResult.Value
	lodResult, _ := LinesOfDocumentation(tokens)
	lod := lodResult.Value

	div := gcd(loc, lod)

	loc = loc / div
	lod = lod / div

	return Result{Metric: "Ratio of LOC to LOD", Value: lod / loc, Body: template.HTML(fmt.Sprintf("The ratio of Lines of Code to Lines of Documentation is %d:%d.", loc, lod))}, nil
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

//Not implemented
func LinesPerFunction(tokens []Token) (Result, error) {
	return Result{Metric: "Lines per Function", Value: 0, Body: template.HTML("not yet implemented")}, nil
}

//Not implemented
func NumberOfFunctionParameters(tokens []Token) (Result, error) {
	return Result{Metric: "Number of Function Parameters", Value: 0, Body: template.HTML("not yet implemented")}, nil
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

//Not implemented
func MethodsPerClass(tokens []Token) (Result, error) {
	return Result{Metric: "Methods per Class", Value: 0, Body: template.HTML("not yet implemented")}, nil
}

//CyclomaticComplexity calculates the complexity of source code
//Not finished
func CyclomaticComplexity(tokens []Token) (Result, error) {
	type function struct {
		name       string
		lines      []string
		complexity int
	}

	functions := make([]function, 0)
	var currFunc function
	nodes := 0
	returns := 0
	braces := 0
	startFunc := false
	for _, token := range tokens {
		if token.Type == "FUNCTION" {
			currFunc = function{name: token.Value, lines: make([]string, 0)}
			nodes = 0
			returns = 0
			braces = 0
			startFunc = true
		} else if token.Type == "IF" || token.Type == "CASE" || token.Type == "FOR" || token.Type == "WHILE" {
			nodes++
		} else if token.Type == "RETURN" {
			returns++
		} else if token.Type == "LEFT_BRACE" {
			braces++
		} else if token.Type == "RIGHT_BRACE" {
			braces--
		}
		currFunc.lines = append(currFunc.lines, token.Value)
		if braces == 0 && !startFunc {
			currFunc.complexity = nodes + 2*returns
		} else if braces == 0 {
			functions = append(functions, currFunc)
		}
	}
	log.Println(functions)
	return Result{Metric: "CyclomaticComplexity", Value: 0, Body: template.HTML(fmt.Sprintf("There are %d nodes", nodes))}, nil
}
