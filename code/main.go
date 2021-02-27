package main

//go application consists of packages and main package is the executable
/*
Go programs are organized into packages.
A package is a collection of source files in the same directory that are compiled together
*/

import (
	"fmt"
	"greetings"
)

// fmt package has string formatting functions

func main() { // main function is the entry point of the application
	printHelloWorld() // simple hello world
	Variables()       // all about variables
	Types()           // about types

	var greetingmsg string=greetings.Hello("anurag")
	fmt.Printf("%s\n",greetingmsg)
	fmt.Scanln()      // wait for user input
}