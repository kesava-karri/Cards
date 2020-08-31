package main

func main() {

	cards := newDeck()

	hand,leftOverCards := deal(cards, 3)

	hand.print()
	leftOverCards.print()
}
