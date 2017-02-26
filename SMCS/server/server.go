package server

import (
	"net/http"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer"
	"html/template"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Config struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
}

func NewConfig(filepath string) *Config {
	config := &Config{}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println(err)
	}
	return config
}

type Handler func(http.ResponseWriter, *http.Request, *Server)

type Server struct {
	Config   *Config
	Handlers map[string]Handler
	Template *template.Template
	Analyzer *analyzer.Analyzer
}

func NewServer(filepath string) *Server {
	config := NewConfig(filepath)

	analyzer := analyzer.NewAnalyzer()

	handlers := make(map[string]Handler)
	handlers["/"] = indexHandler
	handlers["/index.html"] = indexHandler
	handlers["/metrics.html"] = metricsHandler

	tmpl := template.Must(template.New("index.html").New("metrics.html").ParseFiles("assets/tmpl/index.html", "assets/tmpl/metrics.html"))

	return &Server{Config: config, Handlers: handlers, Template: tmpl, Analyzer: analyzer}
}

func (s *Server) Run() {
	for k, v := range s.Handlers {
		http.HandleFunc(k, makeHandler(v, s))
	}

	err := http.ListenAndServe(s.Config.Domain + ":" + s.Config.Port, nil)
	if err != nil {
		log.Println(err)
	}

}
