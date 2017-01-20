//This example package shows that a package can be split up into
//multiple files and it works the same as a single file.
package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Domain   string
	Port     string
	Handlers HandlerMap
}

func New(domain, port string, handlers HandlerMap) *Server {
	return &Server{Domain: domain, Port: port, Handlers: handlers}
}

func (s *Server) Run() {
	//go through the HandlerMap and set all of the handler functions
	for key, val := range s.Handlers {
		http.HandleFunc(key, val)
	}

	//start the server
	//this is a blocking function call, it will only return if there is an error
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.Domain, s.Port), nil)
	if err != nil {
		log.Panicln(err)
	}
}
