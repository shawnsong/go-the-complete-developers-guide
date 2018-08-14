package main

func main() {
	cards := newDeck()

	cards.saveToFile("/tmp/cards")

	newCards := newDeckFromFile("/tmp/cards")

	newCards.shuffle()
	newCards.print()
}
