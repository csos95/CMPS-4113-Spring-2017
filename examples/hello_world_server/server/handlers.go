//This example package shows that a package can be split up into
//multiple files and it works the same as a single file.
package server

import (
	"fmt"
	"net/http"
)

//create HandlerMap type to reduce clutter where it is used
//this is similar to typedef in c
type HandlerMap map[string]func(http.ResponseWriter, *http.Request)

//A handler function that can be used with http.HandleFunc
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//write a string to a io.Writer using default formatting
	fmt.Fprintln(w, "Go to /hello or /bye")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func ByeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye!")
}
