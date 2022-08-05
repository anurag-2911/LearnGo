package main

import "fmt"

// create a new type of desk which has slice of strings

type deck []string

//receiver function
func (cards deck) printDeck() {
	for _, card := range cards {
		fmt.Println(card)

	}

}

// creates new deck

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "hearts", "club"}

	cardValues := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuits {

		for _, cardvalue := range cardValues {
			cards = append(cards, suit+" and "+cardvalue)
		}
	}
	return cards
}
