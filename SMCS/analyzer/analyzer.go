package analyzer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

type Result struct {
	Metric string
	Value  int
	Body   template.HTML
}

type Analysis struct {
	Language *Language
	Source   string
	Results  []Result
}

type Language struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Homepage      string   `json:"homepage"`
	Extensions    []string `json:"extensions"`
	Keywords      []string `json:"keywords"`
	Operators     []string `json:"operators"`
	LineComments  []string `json:"linecomments"`
	BlockComments []string `json:"blockcomments"`
}

func NewLanguage(filepath string) *Language {
	lang := &Language{}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(file, &lang)
	if err != nil {
		log.Println(err)
	}
	return lang
}

type Analyzer struct {
	Languages map[string]*Language
	Metrics   map[string]Metric
}

func NewAnalyzer() *Analyzer {
	languages := make(map[string]*Language)

	languages["c"] = &Language{Name: "c", Extensions: []string{".c", ".h"}}
	languages["cpp"] = &Language{Name: "cpp", Extensions: []string{".cpp", ".hpp"}}
	languages["java"] = &Language{Name: "java", Extensions: []string{".java"}}

	metrics := map[string]Metric{}

	metrics["Lines of Code"] = LinesOfCode
	metrics["Lines of Documentation"] = LinesOfDocumentation
	metrics["Ratio of LOC to LOD"] = RatioOfLOCToLOD
	metrics["Blank Lines"] = BlankLines
	metrics["Total Lines"] = TotalLines

	metrics["Number of Functions"] = NumberOfFunctions
	//metrics["Lines per Function"] = LinesPerFunction
	//metrics["Number of Function Parameters"] = NumberOfFunctionParameters
	metrics["Number of Classes"] = NumberOfClasses
	//metrics["Methods per Class"] = MethodsPerClass

	//metrics["Cyclomatic Complexity"] = CyclomaticComplexity

	return &Analyzer{Languages: languages, Metrics: metrics}
}

func (a *Analyzer) AddLanguage(language *Language) {
	a.Languages[language.Name] = language
}

func (a *Analyzer) Extensions() map[string][]string {
	extensions := make(map[string][]string)
	for _, language := range a.Languages {
		extensions[language.Name] = language.Extensions
	}
	return extensions
}

func (a *Analyzer) Analyze(language, source string, metrics []string) Analysis {
	analysis := Analysis{Language: a.Languages[language], Source: source}

	tokens := Tokenize(language, analysis.Source)

	for _, metric := range metrics {
		result, err := a.Metrics[metric](tokens)
		if err != nil {
			analysis.Results = append(analysis.Results, Result{Metric: metric, Body: template.HTML(fmt.Sprintf("Error: %v", err))})
			log.Println(err)
		} else {
			analysis.Results = append(analysis.Results, result)
		}
	}
	return analysis
}
