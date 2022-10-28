package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck()

	hand, remainingDeck := cards.deal(5)
	fmt.Println("\nHand:")
	hand.print()
	fmt.Println("\nRemaining:")
	remainingDeck.print()

	/* cards.saveToFile("My Deck")

	greeting := "Hello"
	fmt.Println([]byte(greeting))

	fmt.Println("\nNew Deck of Cards\n")
	newCards := newDeckFromFile("My Deck")
	newCards.print() */
	fmt.Println("\nRemaining Deck Shuffled:")
	remainingDeck.shuffle().print()

}
