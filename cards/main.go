package main

import (
	"fmt"
	"strconv"
)

func main() {
	callBasics()
}

func callBasics() {
	fmt.Print(" hello world ")
	var card string = " variable declaration "
	fmt.Print(card)
	var1 := " another name "
	var1 = " yes "
	fmt.Print(var1)
	fmt.Print(name() + " " + strconv.Itoa(age()))
	var names []string = append(getNames(), "khalkho")
	fmt.Print(names)
	printNames(names)
}
func name() string {
	return " anurag "
}
func age() int {
	return 1
}

func getNames() []string {
	var nms = []string{"topno", "kunar", "benjamin"}
	return nms
}

func printNames(names []string) {
	for i, name := range names {
		fmt.Print(i, name)
		fmt.Print(" ")
	}
}
