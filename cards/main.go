/* card game */

package main

func main() {
	//initialize(create) and assign the variable with :=
	// var card string = "Ace of Spades"
	// cards := newDeck()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//cards.saveToFile("myCards")
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//println(remainingCards.toString())
}
