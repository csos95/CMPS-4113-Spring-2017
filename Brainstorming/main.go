package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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

type Metric func(source string) (Result, error)

func LinesOfCode(source string) (Result, error) {
	result := Result{Metric: "Lines of Code"}
	lines := strings.Split(source, "\n")
	lineCount := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !(len(line) > 1 && line[:2] == "//") {
			lineCount++
		}
	}
	result.Body = template.HTML(fmt.Sprintf("<div id=\"loc\" class=\"metric\"><p>There are %d lines of code.</p></div>", lineCount))
	return result, nil
}

func LinesOfDocumentation(source string) (Result, error) {
	result := Result{Metric: "Lines of Documentation"}
	lines := strings.Split(source, "\n")
	lineCount := 0
	for _, line := range lines {
		if strings.Contains(line, "//") {
			lineCount++
		}
	}
	result.Body = template.HTML(fmt.Sprintf("<div id=\"lod\" class=\"metric\"><p>There are %d lines of documentation.</p></div>", lineCount))
	return result, nil
}

func BlankLines(source string) (Result, error) {
	result := Result{Metric: "Blank Lines"}
	lines := strings.Split(source, "\n")
	lineCount := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			lineCount++
		}
	}
	result.Body = template.HTML(fmt.Sprintf("<div id=\"bl\" class=\"metric\"><p>There are %d blank lines.</p></div>", lineCount))
	return result, nil
}

type Analyzer struct {
	Lang    *Language
	Metrics []Metric
	Source  string
}

func NewAnalyzer(lang *Language, metrics []Metric) *Analyzer {
	return &Analyzer{Lang: lang, Metrics: metrics}
}

func (a *Analyzer) Analyze(source string) Analysis {
	analysis := Analysis{Source: source, Results: make([]Result, len(a.Metrics))}
	var err error
	for i, metric := range a.Metrics {
		analysis.Results[i], err = metric(source)
		if err != nil {
			log.Println(err)
		}
	}
	return analysis
}

type Analysis struct {
	Source  string
	Results []Result
}

type Result struct {
	Metric string
	Body   template.HTML
}

func main() {
	l := NewLanguage("language_example.json")
	fmt.Println(l)

	interpreter := NewAnalyzer(l, []Metric{LinesOfCode, LinesOfDocumentation, BlankLines})

	var source string

	tmpl := template.Must(template.New("metrics.html").New("index.html").ParseFiles("assets/tmpl/metrics.html", "assets/tmpl/index.html"))
	http.HandleFunc("/metrics/", func(w http.ResponseWriter, r *http.Request) {
		analysis := interpreter.Analyze(source)
		tmpl.ExecuteTemplate(w, "metrics.html", analysis)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})
	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseMultipartForm(32 << 20)
			file, _, err := r.FormFile("uploadfile")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			buf := bytes.NewBuffer(nil)
			io.Copy(buf, file)
			source = string(buf.Bytes())
			fmt.Println("Source:", source)
		}
		http.Redirect(w, r, "http://127.0.0.1:8080/metrics/", http.StatusFound)
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	assetfs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetfs))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/assets/img/favicon.ico", http.StatusMovedPermanently)
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Println(err)
	}
}
