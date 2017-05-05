package server

import (
	"bytes"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer"
	"io"
	"log"
	"net/http"
)

//Page holds the data needed to create a html page
type Page struct {
	Config     *Config
	Extensions map[string][]string
	Source     string
	Analysis   analyzer.Analysis
	Languages  map[string]*analyzer.Language
	Metrics    map[string]analyzer.Metric
}

//makeHandler takes a handler with a server parameter and returns a function matching the http.Handler interface
func makeHandler(fn func(http.ResponseWriter, *http.Request, *Server), s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s)
	}
}

//indexHandler handles the source code selection page
func indexHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	if r.Method == "GET" {
		//create the source upload page with available languages and metrics
		page := &Page{Config: s.Config, Extensions: s.Analyzer.Extensions(), Metrics: s.Analyzer.Metrics, Languages: s.Analyzer.Languages}

		//display the source upload page
		s.Template.ExecuteTemplate(w, "index.html", page)
	}
}

//metricsHandler handlers the metrics page
func metricsHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	if r.Method == "POST" {
		//parse the parameters and file sent in the post request
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()

		buff := bytes.NewBuffer(nil)
		io.Copy(buff, file)
		source := string(buff.Bytes())

		//get selected language
		language := r.Form["language"][0]

		//get selected metrics
		metrics := make([]string, 0)
		selectedMetrics := r.Form["metric"]
		for _, metric := range selectedMetrics {
			if _, ok := s.Analyzer.Metrics[metric]; ok {
				metrics = append(metrics, metric)
			}
		}

		//analyze the code with selected metrics
		page := &Page{Config: s.Config, Source: source, Languages: s.Analyzer.Languages,
			Analysis: s.Analyzer.Analyze(language, source, metrics)}

		//display the results page
		s.Template.ExecuteTemplate(w, "metrics.html", page)
	} else if r.Method == "GET" {
		http.Redirect(w, r, "http://"+s.Config.Domain+":"+s.Config.Port+"/", http.StatusFound)
	}
}
