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

//Config holds the application configuration options
type Config struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
	Mode   string `json:"mode"`
}

//LoadConfig loads the config from the specified path
func LoadConfig(filepath string) (*Config, error) {
	config := &Config{}

	//attempt to read config file
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		//create a default config if that file doesn't exist
		return DefaultConfig(filepath)
	}

	//parse the config into the config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

//DefaultConfig creates, writes, and returns the default config
func DefaultConfig(filepath string) (*Config, error) {
	//the default config
	config := &Config{Domain: "0.0.0.0", Port: "8080", Mode: "server"}

	//marshal it into a json
	js, _ := json.Marshal(config)

	//get the parent filepath of the config file
	parts := strings.Split(filepath, "/")
	path := strings.Join(parts[:len(parts)-1], "/")

	//make sure that filepath exists before trying to put a file there
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	//write the default config
	err := ioutil.WriteFile(filepath, js, os.ModePerm)
	if err != nil {
		return config, err
	}

	return config, nil
}

//Handler is a handler function that accepts a server
type Handler func(http.ResponseWriter, *http.Request, *Server)

//Server contains the application config, webpage templates, and the analyzer
type Server struct {
	Config   *Config
	Template *template.Template
	Analyzer *analyzer.Analyzer
}

//NewServer creates a new server with the specified config file
func NewServer(filepath string) *Server {
	//load the config
	config, err := LoadConfig(filepath)
	if err != nil {
		log.Println(err)
	}

	//create a new analyzer
	analyzer := analyzer.NewAnalyzer()

	//create the template with the two webpage templates
	tmpl := template.New("")
	files := []string{"index", "metrics"}

	//load them from the bindata_assetfs
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

//Run executes the webserver
func (s *Server) Run() error {
	//create a new mux router
	router := mux.NewRouter()

	//setup the assets handlers
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(assetFS())))
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/assets/img/favicon.ico", http.StatusMovedPermanently)
	})

	//setup the template handlers
	router.HandleFunc("/", makeHandler(indexHandler, s)).Methods("GET")
	router.HandleFunc("/index.html", makeHandler(indexHandler, s)).Methods("GET")
	router.HandleFunc("/metrics.html", makeHandler(metricsHandler, s)).Methods("POST")

	//if running in standalone, launch the users default browser
	if s.Config.Mode == "standalone" {
		go func() {
			time.Sleep(time.Second * 2)
			openBrowser("http://" + s.Config.Domain + ":" + s.Config.Port + "/")
		}()
	}

	//create a server using mux router
	srv := &http.Server{
		Handler: router,
		Addr:    s.Config.Domain + ":" + s.Config.Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//start the server
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
