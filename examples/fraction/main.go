//This example shows how to import and use a user written package.
//The fraction package is imported using the gopath relative path.
//Because this fraction package is nested four folders from the repo
//root, the import path is very long.
//Normally it would be something like "github.com/csos95/fraction".
package main

//import the fraction class which is accessed with fraction.[method]
import (
	"github.com/csos95/CMPS-4113-Spring-2017/examples/fraction/fraction"
	"fmt"
)

func main() {
	//create two fractions and print them with Println and Printf
	foo := fraction.New(1, 2)
	fmt.Println("Foo =", foo)
	bar := fraction.New(3, 5)
	fmt.Printf("Bar = %v\n", bar)

	//print the sum of the two using both add methods.
	fmt.Println("Foo + Bar =", fraction.Addr(foo, bar))
	foo.Add(bar)
	fmt.Printf("Foo += Bar\nFoo = %v\n", foo)
}
