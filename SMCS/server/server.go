package server

import (
	"encoding/json"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"os/exec"
	"time"
)

type Config struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
	Mode string `json:"mode"`
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

	assetfs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetfs))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/assets/img/favicon.ico", http.StatusMovedPermanently)
	})

	if s.Config.Mode == "standalone" {
		go func() {
			time.Sleep(time.Second * 2)
			openBrowser("http://"+s.Config.Domain+":"+s.Config.Port+"/")
		}()
	}

	err := http.ListenAndServe(s.Config.Domain+":"+s.Config.Port, nil)
	if err != nil {
		log.Println(err)
	}

}

//openBrowser opens the default user browser with the specified url
//taken from github.com/rodzzlessa24/openbrowser.go
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
