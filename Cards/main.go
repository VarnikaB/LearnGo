package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := []string{newCard(), "Two"}
	cards = append(cards, "Three")
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "One of Diamonds"
}