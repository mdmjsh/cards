package main

func main() {
	// eqiv. to `var card string = "Ace of Spades"`
	// card := "Ace of Spades"
	// card := newCard()

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// hand, remainingDeck := deal(cards, 5)
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
