package main

func main() {
	// long form method
	// could have omitted the string
	//var card string = "Ace of Spades"

	// only used when defining a new variable
	// initialization statment
	//card := "Ace of Spades"

	// var card = "Ace of Spades"

	// this creates a slice
	// which is an array without a defined size
	// like a list basically
	// cards := newDeck()
	// fmt.Println((cards.toString()))
	// cards.saveToFile("my_cards")

	// iterate over a slice
	// can call this because cards is of type deck
	// cards.print()
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}