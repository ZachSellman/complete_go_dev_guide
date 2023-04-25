package main

type deck []string

var suits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
var values = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

func buildNewDeck() deck {
	d := deck{}
	for _, suit := range suits {
		for _, value := range values {
			d = append(d, value+" of "+suit)
		}
	}
	return d
}
