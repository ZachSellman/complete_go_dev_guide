package main

import (
	"fmt"
)

func main() {
	deck := newDeck()
	// card := getRandom(deck)
	fmt.Println("deck: ", deck)
}

func newDeck() []int {
	var deck []int
	for i := 1; i < 53; i++ {
		deck = append(deck, i)
	}
	return deck
}
