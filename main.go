package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("cards")
	//cards := newDeckFromFile("cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
