package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	deal(cards, 5)
	cards.print()
	cards.saveToFile("my_cards")
	newDeck := newDeckFromFile("my_cards")
	newDeck.shuffle()
	newDeck.print()
}
