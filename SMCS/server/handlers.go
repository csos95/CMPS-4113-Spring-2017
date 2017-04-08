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
//<input type="checkbox" name="loc" value="Lines of Code">Lines of Code<br>
//<input type="checkbox" name="lod" value="Lines of Documentation">Lines of Documentation<br>
//<input type="checkbox" name="loc/lod" value="Ratio of LOC to LOD">Ratio of LOC to LOD<br>
//<input type="checkbox" name="bl" value="Blank Lines">Blank Lines<br>
//<input type="checkbox" name="tl" value="Total Lines">Total Lines<br>
//
//<input type="checkbox" name="nf" value="Number of Functions">Number of Functions<br>
//<input type="checkbox" name="nfp" value="Number of Function Parameters">Number of Function Parameters<br>
//<input type="checkbox" name="nc" value="Number of Classes">Number of Classes<br>
//<input type="checkbox" name="mc" value="Methods per Class">Methods per Class<br>
//<input type="checkbox" name="lf" value="Lines per Function">Lines per Function<br>
//
//<input type="checkbox" name="cc" value="Cyclomatic Complexity">Cyclomatic Complexity<br>
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
