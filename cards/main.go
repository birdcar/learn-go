package main

import (
	"flag"
	"fmt"
)

var cards deck
var inFile string
var outFile string
var shuffle bool

func main() {
	flag.StringVar(
		&inFile,
		"in",
		"",
		"[-in]: an optionally specified path to an input file containing a deck of cards",
	)
	flag.StringVar(
		&outFile,
		"out",
		"",
		"[-out]: an optionally specified path to a file to save the deck to",
	)
	flag.BoolVar(
		&shuffle,
		"shuffle",
		false,
		"[-shuffle]: Optionally ask the program to shuffle the deck",
	)
	flag.Parse()

	// If inFile has been specified, load from a file, otherwise create a new deck
	if len(inFile) > 0 {
		cards = newDeckFromFile(inFile)
	} else {
		cards = newDeck()
	}

	// Shuffle the deck if the shuffle command line arg has been provided
	if shuffle {
		cards.shuffle()
	}

	if len(outFile) > 0 {
		cards.saveToFile(outFile)
	} else {
		fmt.Println(cards)
	}
}
