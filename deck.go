package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
// n.b OO equiv of subclassing/extending a []string
type deck []string

// Receiver function - takes a deck and prints it
// OO equiv of a method on a class
// by convention, the receiver is a one or two letter abbreviation of the type
func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

func newDeck() deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filepath string) error {
	return os.WriteFile(filepath, []byte(d.toString()), 0666)
}

func newDeckFromFile(filepath string) deck {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bytes), ","))
}

func (d deck) shuffle() {
	deckLen := len(d) - 1
	// seed a RNG using the unix time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(deckLen)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
