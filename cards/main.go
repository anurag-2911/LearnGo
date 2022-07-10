package main

import "fmt"

func main() {
	var card string = "first card" // variable declaration
	fmt.Println(card);

	name:="anurag" // variable declaration
	fmt.Println(name);

	var nm string = getName();
	fmt.Println(nm);

	
	// array | slice

	cards:= deck{"card1","card2","card3"}

	// fmt.Println(cards)

	cards = append(cards, "card3")

	// fmt.Println(cards)
    // for loop
	// for i, v := range cards {
	// 	fmt.Println(i,v);
	// }

	cards.print()
}

func getName() string {
	return "anurag";
}