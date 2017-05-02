package server

import (
	"encoding/json"
	"fmt"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Config struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
	Mode   string `json:"mode"`
}

func NewConfig(filepath string) (*Config, error) {
	config := &Config{}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return DefaultConfig(filepath)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func DefaultConfig(filepath string) (*Config, error) {
	config := &Config{Domain: "0.0.0.0", Port: "8080", Mode: "server"}

	js, _ := json.Marshal(config)

	parts := strings.Split(filepath, "/")
	path := strings.Join(parts[:len(parts)-1], "/")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	err := ioutil.WriteFile(filepath, js, os.ModePerm)
	if err != nil {
		return config, err
	}
	return config, nil
}

type Handler func(http.ResponseWriter, *http.Request, *Server)

type Server struct {
	Config   *Config
	Template *template.Template
	Analyzer *analyzer.Analyzer
}

func NewServer(filepath string) *Server {
	config, err := NewConfig(filepath)
	if err != nil {
		log.Println(err)
	}

	analyzer := analyzer.NewAnalyzer()

	tmpl := template.New("")
	files := []string{"index", "metrics"}

	for _, file := range files {
		data, err := Asset(fmt.Sprintf("assets/tmpl/%s.html", file))
		if err != nil {
			log.Println(err)
		}
		tmpl = template.Must(tmpl.New(fmt.Sprintf("%s.html", file)).Parse(string(data)))
	}

	//tmpl := template.Must(template.New("index.html").New("metrics.html").ParseFiles("assets/tmpl/index.html", "assets/tmpl/metrics.html"))

	return &Server{Config: config, Template: tmpl, Analyzer: analyzer}
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(assetFS())))
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/assets/img/favicon.ico", http.StatusMovedPermanently)
	})

	router.HandleFunc("/", makeHandler(indexHandler, s)).Methods("GET")
	router.HandleFunc("/index.html", makeHandler(indexHandler, s)).Methods("GET")
	router.HandleFunc("/metrics.html", makeHandler(metricsHandler, s)).Methods("POST")

	if s.Config.Mode == "standalone" {
		go func() {
			time.Sleep(time.Second * 2)
			openBrowser("http://" + s.Config.Domain + ":" + s.Config.Port + "/")
		}()
	}

	srv := &http.Server{
		Handler: router,
		Addr:    s.Config.Domain + ":" + s.Config.Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
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
