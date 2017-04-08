package server

import (
	"bytes"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer"
	"io"
	"log"
	"net/http"
)

type Page struct {
	Config     *Config
	Extensions map[string][]string
	Source     string
	Analysis   analyzer.Analysis
	Language   *analyzer.Language
	Metrics    map[string]analyzer.Metric
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *Server), s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	if r.Method == "GET" {
		page := &Page{Config: s.Config, Extensions: s.Analyzer.Extensions(), Metrics: s.Analyzer.Metrics}
		s.Template.ExecuteTemplate(w, "index.html", page)
	}
}

func metricsHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	if r.Method == "POST" {
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

		metrics := make([]string, 0)
		log.Println(r.Form)
		selectedMetrics := r.Form["metric"]
		for _, metric := range selectedMetrics {
			log.Println("checking for", metric)
			if _, ok := s.Analyzer.Metrics[metric]; ok {
				log.Println("ok")
				metrics = append(metrics, metric)
			}
		}

		log.Println(metrics)

		page := &Page{Config: s.Config, Source: source, Language: s.Analyzer.Languages["c"],
			Analysis: s.Analyzer.Analyze("c", source, metrics)}

		s.Template.ExecuteTemplate(w, "metrics.html", page)
	} else if r.Method == "GET" {
		http.Redirect(w, r, "http://"+s.Config.Domain+":"+s.Config.Port+"/", http.StatusFound)
	}
}
