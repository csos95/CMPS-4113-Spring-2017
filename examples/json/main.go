//This example show how to read and parse a json file
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//Example holds a string message
//The Message variable has a json tag on it
//This tag is used by the json.Unmarshal method
//when parsing the json data
type Example struct {
	Message string `json:"message"`
}

func main() {
	//Read the json file.
	file, err := ioutil.ReadFile("example.json")
	//If there were any errors reading the file, log the stack trace and exit the program.
	if err != nil {
		log.Panicln(err)
	}
	//In Go, the error handling pattern is
	//value, err := function()
	//if err != nil {
	//    handle error
	//}
	//Usually you will not want to panic because of an error, but I do so here because the
	//only point of this example is to read, parse, and print a message from a json file

	//create a Example variable to store the parsed data in
	var ex Example

	err = json.Unmarshal(file, &ex)
	//If there were any errors parsing the json, log the stack trace and exit the program.
	if err != nil {
		log.Panicln(err)
	}

	//print the parsed message
	fmt.Println(ex.Message)
}
