package main

import "fmt"

// Create a new type deck which is slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := deck{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" Of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
