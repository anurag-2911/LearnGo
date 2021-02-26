package main

//go application consists of packages and main package is the executable
/*
Go programs are organized into packages.
A package is a collection of source files in the same directory that are compiled together
*/

import (
	"fmt"
)

// fmt package has string formatting functions

func main() { // main function is the entry point of the application
	PrintHelloWorld()
	One()
	fmt.Scanln() // wait for user input
}
