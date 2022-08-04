package main

import "fmt"

// create a new type of desk
// which has slice of strings

type deck []string

//receiver function
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)

	}

}
