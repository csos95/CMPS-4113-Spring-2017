//This example shows how to use a fileserver to serve files in a directory,
//how to create and use a simple template file, and how to use a makeHandler function.
//For more information on Go templates visit https://golang.org/pkg/text/template/
package main

import (
	"html/template"
	"log"
	"net/http"
)

//Page represents a simple page with just a title and body
type Page struct {
	Title string
	Body  string
}

type Handler func(http.ResponseWriter, *http.Request, *template.Template)

//the idiomatic way in go to write a handler that requires more than
//the two http.Handler parameters is to use a makeHandler function
//this function will receive the handler function you wish to use,
//the extra parameters it needs, and returns a http.HandlerFunc
func makeHandler(h Handler, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL, r.UserAgent())
		//this function has access to the variables passed to makeHandler.
		//we are therefore able to call the handler function with the
		//additional template parameter from inside a http.Handler function
		h(w, r, tmpl)
	}
}

//indexHander uses a template to generate a html page with a title and body
//the member data of the object passed as the second parameter is accessible within
//the template file.
//Template features are signified by "{{}}".
//To access the Title variable, you would use {{.Title}} amd the value of Title
//will replace it when executed.
func indexHandler(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	tmpl.Execute(w, &Page{Title: "This is the title.", Body: "This is the body."})
}

func main() {
	//parse index.html into a template and compile it
	//template.Must requires that the template it is given compiles correctly
	//if it does not, it will panic
	indexTmpl := template.Must(template.New("index.html").ParseFiles("assets/tmpl/index.html"))
	http.HandleFunc("/", makeHandler(indexHandler, indexTmpl))

	//create a fileserver (type http.Handler) for assets
	assetfs := http.FileServer(http.Dir("assets"))
	//use it to handle urls starting with "/assets/"
	http.Handle("/assets/", http.StripPrefix("/assets/", assetfs))

	//the browser will attempt to access /favicon.ico
	//redirect it to /assets/img/favicon.ico so it uses the assetfs
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/assets/img/favicon.ico", http.StatusMovedPermanently)
	})

	//start the server listening at 127.0.0.1:8080
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Panicln(err)
	}
}
