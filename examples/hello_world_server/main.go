//This example shows that a package split into multiplefiles works
//the same as one in a single file.
//This example creates a server that serves up one of three
//messages depending on the url requested.
package main

import (
	"github.com/csos95/CMPS-4113-Spring-2017/examples/hello_world_server/server"
)

func main() {
	//create a new HandlerMap (Defined as map[string]func(http.ResponseWriter, *http.Request))
	handlers := make(server.HandlerMap)

	//set the handler that will be used for three url paths
	//if no other paths match, "/" will be used
	handlers["/"] = server.IndexHandler
	handlers["/hello"] = server.HelloHandler
	handlers["/bye"] = server.ByeHandler

	//create a new server that is accessable at 127.0.0.1:8080
	s := server.New("127.0.0.1", "8080", handlers)

	//run the server
	s.Run()
}
