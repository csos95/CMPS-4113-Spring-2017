package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Language struct {
	Name          string   `json:"language"`
	Description   string   `json:"description"`
	Homepage      string   `json:"homepage"`
	Extensions    []string `json:"extensions"`
	Keywords      []string `json:"keywords"`
	Operators     []string `json:"operators"`
	LineComments  []string `json:"linecomments"`
	BlockComments []string `json:"blockcomments"`
}

func NewLanguage(filepath string) *Language {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println(err)
		return nil
	}

	l := &Language{}
	err = json.Unmarshal(file, &l)
	if err != nil {
		log.Println(err)
		return nil
	}
	return l
}

func (l *Language) String() string {
	return fmt.Sprintf("Name: %s\nDescription: %s\nHomepage: %v\nExtensions: %v", l.Name, l.Description, l.Homepage, l.Extensions)
}

type Metric func(source string) (template.HTML, error)

func LinesOfCode(source string) (template.HTML, error) {
	lines := strings.Split(source, "\n")
	lineCount := 0
	for _, line := range lines {
		if !strings.Contains(line, "//") {
			lineCount++
		}
	}
	html := template.HTML(fmt.Sprintf("<div id=\"loc\" class=\"metric\"><p>There are %d lines of code.</p></div>", lineCount))
	return html, nil
}

func LinesOfDocumentation(source string) (template.HTML, error) {
	lines := strings.Split(source, "\n")
	lineCount := 0
	for _, line := range lines {
		if strings.Contains(line, "//") {
			lineCount++
		}
	}
	html := template.HTML(fmt.Sprintf("<div id=\"loc\" class=\"metric\"><p>There are %d lines of documentation.</p></div>", lineCount))
	return html, nil
}

type Analyzer struct {
	Lang    *Language
	Metrics []Metric
}

func NewAnalyzer(lang *Language, metrics []Metric) *Analyzer {
	return &Analyzer{Lang: lang, Metrics: metrics}
}

func (a *Analyzer) Analyze(source string) []template.HTML {
	results := make([]template.HTML, len(a.Metrics))
	var err error
	for i, metric := range a.Metrics {
		results[i], err = metric(source)
		if err != nil {
			log.Println(err)
		}
	}
	return results
}

func main() {
	l := NewLanguage("language_example.json")
	fmt.Println(l)

	interpreter := NewAnalyzer(l, []Metric{LinesOfCode, LinesOfDocumentation})

	source := `package main
	import "fmt"
	//The main function is required as an entrypoint if this is a application and not a library
	func main() {
		fmt.Println("Hello, World!")
	}`

	results := interpreter.Analyze(source)
	for _, result := range results {
		fmt.Println(result)
	}
	tmpl := template.Must(template.New("metrics.html").ParseFiles("metrics.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "metrics.html", results)
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Println(err)
	}
}
