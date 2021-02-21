package main

import (
	"fmt"
	"strconv"
)

//variables at package level
var (
	name    string = "anurag"
	count   int    = 10
	message string = "learn"
)

// I is a global scope variable
var I int = 100

func main() {
	//variable declared in function scope
	var i int = 42 // declared and assigned
	j := 43        // compiler detects the variable type here
	var l int      // just declared
	l = 99
	fmt.Printf("%d,%d,%d\n",i, j, l)
	fmt.Println("hello world")
	fmt.Println(message)
	fmt.Println(I);
	conversion();
}

func conversion() {
	var i int =50
	fmt.Printf("%v , %T\n",i,i);
	var j float32
	j=float32(i)
	fmt.Printf("%v , %T\n",j,j);
	var s=strconv.Itoa(i) // for string conversion
	fmt.Println(s);
}
