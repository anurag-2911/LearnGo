package main

import (
	"fmt"
)

func main() {
	callBasics()
}

func callBasics() {
	fmt.Print(" hello world!! ")
	
	cards := newDeck()
	
	cards.printDeck()	
	
}


