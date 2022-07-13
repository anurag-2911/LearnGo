package main

import "fmt"

type deck []string

//receiver function
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i,card)

	}

}
