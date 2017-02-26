package server

import (
	"net/http"
	"log"
	"bytes"
	"io"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer"
)

type Page struct {
	Config *Config
	Extensions map[string][]string
	Analysis analyzer.Analysis
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

		page := &Page{Config: s.Config, Analysis: s.Analyzer.Analyze("c++", source, []string{"Lines of Code", "Lines of Documentation"})}

		s.Template.ExecuteTemplate(w, "metrics.html", page)
	} else if r.Method == "GET" {
		s.Template.ExecuteTemplate(w, "metrics.html", s)
	}
}
