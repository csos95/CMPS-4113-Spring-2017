package analyzer

import (
	"html/template"
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
)

type Result struct {
	Metric string
	Body template.HTML
}

type Analysis struct {
	Source string
	Results []Result
}

type Language struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Homepage string `json:"homepage"`
	Extensions []string `json:"extensions"`
	Keywords []string `json:"keywords"`
	Operators []string `json:"operators"`
	LineComments []string `json:"linecomments"`
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
	Metrics map[string]Metric
}

func NewAnalyzer() *Analyzer {
	languages := make(map[string]*Language)

	languages["c++"] = &Language{Name: "c++", Extensions: []string{".cpp", ".h", ".hpp"}}
	languages["java"] = &Language{Name: "java", Extensions: []string{".java"}}
	languages["go"] = &Language{Name: "go", Extensions: []string{".go"}}

	metrics := map[string]Metric{}

	metrics["Lines of Code"] = LinesOfCode
	metrics["Lines of Documentation"] = LinesOfDocumentation

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
	analysis := Analysis{Source: source}

	tokens := Tokenize(language, analysis.Source)

	for _, metric := range metrics {
		log.Println(metric)
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
