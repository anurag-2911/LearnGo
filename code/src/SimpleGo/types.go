package main

import (
	"fmt"
)

// Types example
func Types(){
	fmt.Printf("%s","about go data types\n")
	booleanType()
}

func booleanType(){
	var b bool

	fmt.Printf("%v,%T\n",b,b)
}