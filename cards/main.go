package main

import 
(
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(" hello world ")
	var card string = " variable declaration "
	fmt.Print(card)
	var1 := " another name "
	var1 = " yes "
	fmt.Print(var1)
	fmt.Print(names() + " " + strconv.Itoa(age()))
	
}
func names() string {
	return " anurag ";
}
func age() int {
	return 1;
}