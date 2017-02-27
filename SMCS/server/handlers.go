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
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *Server), s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, s *Server) {
	if r.Method == "GET" {
		page := &Page{Config: s.Config, Extensions: s.Analyzer.Extensions()}
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

		page := &Page{Config: s.Config,
			Analysis: s.Analyzer.Analyze("c", source, []string{"Lines of Code", "Lines of Documentation", "Number of Functions"}),
			Source:   source, Language: s.Analyzer.Languages["c"]}

		s.Template.ExecuteTemplate(w, "metrics.html", page)
	} else if r.Method == "GET" {
		s.Template.ExecuteTemplate(w, "metrics.html", s)
	}
}
