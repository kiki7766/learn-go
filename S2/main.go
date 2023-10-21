package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	card := OutCard()
	fmt.Println(card)
}

func OutCard() string {
	return "Five of Diamonds"
}
