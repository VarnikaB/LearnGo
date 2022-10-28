package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck()

	hand, remainingDeck := cards.deal(5)
	fmt.Println("\nHand\n")
	hand.print()
	fmt.Println("\nRemaining\n")
	remainingDeck.print()

	cards.saveToFile("My Deck")
	/* greeting := "Hello"
	fmt.Println([]byte(greeting)) */
	fmt.Println("\nNew Deck of Cards\n")
	newCards := newDeckFromFile("My Deck")
	newCards.print()
}
