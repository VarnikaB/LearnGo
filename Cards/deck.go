package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (cards deck) deal(handsize int) (deck, deck) {
	return cards[:handsize], cards[handsize:]
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), "\n")
}

func (cards deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(cards.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
		//return newDeck()
	}
	s := strings.Split(string(bs), "/n")
	return deck(s)
}
