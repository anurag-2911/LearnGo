package main

import "fmt"

//variables at package level
var (
	name string ="anurag"
	count int =10
	message string ="learn"
)

func main(){
	//variable declared in function scope
	var i int =42 // declared and assigned
	j:=43 // compiler detects the variable type here
	var l int // just declared
	l=99
	fmt.Println(i,j,l);
	fmt.Println("hello world")
	fmt.Println(message);
}