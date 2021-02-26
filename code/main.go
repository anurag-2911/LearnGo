package main

//main is the entry point of the application,go application consists of packages and main package is the executable
import (
	"fmt"
)

// fmt package has string formatting functions

func main() {
	printHelloWorld()
}

func printHelloWorld() {
	fmt.Printf("%s\n", "hello world")
}
