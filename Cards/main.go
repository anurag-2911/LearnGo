package main

import "fmt"

func main() {

	//long variable declaration
	var input string = "learning something new!"

	// short variable declaration
	helloVar := callHello()

	//slice of variable length string array
	arrayVariable := []string{callHello(), callHello() + "test1", callHello() + "test2"}

	// adding one more value in slice
	arrayVariable = append(arrayVariable, callHello()+"test3")

	fmt.Println((arrayVariable))
	fmt.Println(input)
	fmt.Println(helloVar)

	//iterate over a slice

	for index, currentVariable := range arrayVariable {
		fmt.Print(index, currentVariable)
	}

}

func callHello() string {
	return "hello world!"
}
