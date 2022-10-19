package main

func main() {
	//var card string = "Ace of Spades"
	cards := deck{newCard(), "Two"}
	cards = append(cards, "Three")

	cards.print()

}

func newCard() string {
	return "One of Diamonds"
}
