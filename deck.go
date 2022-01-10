package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

// create a new function which prints out each individual card in a deck

// this uses a receiver
// any variable of type deck, now gets access to the print function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// not a receiver method
func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}
	
	// the underscore is that we understand there is a variable here
	// but we don't want to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// this is not a function OF type of deck
// it is just defined here
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// turns a deck into a string
// as a receiver so we can call deck.toString
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// this uses default permissions
// which means anyone can read or write this file
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	// gets the comma separated byte slice
	// and if there was an error msg
	// "nil" means no value for error
	byteSlice, err := os.ReadFile(fileName)

	// this means we had an error
	if err != nil {
		// option 1: log the error and return a call to new deck
		// option 2: log the error and quit the program
		
		// choosing option 2
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// need to convert the list from a []byte
	// convert to string
	// split to []string on comma separator
	// convert to deck
	s := strings.Split(string(byteSlice), ",")
	// can convert this to deck because deck is really just []string
	return deck(s)
}

func (d deck) shuffle() {
	// this is like our seed value
	// look at docs to see why we did this
	// time.now.UnixNano gets a new seed value every time we run``
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// truly random
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// one line to swap elements in a slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}

	// not truly random
	// for i := range d {
	// 	newPosition := rand.Intn(len(d) - 1)

	// 	// one line to swap elements in a slice
	// 	d[i], d[newPosition] = d[newPosition], d[i]
	// }
}