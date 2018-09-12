package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
	create a new type of 'deck'
	which is a slice of strings
*/

type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heards", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	// return value
	return cards
}

// receiver of a function. deck_instance.print() method. 'd' is like a 'self' in python.
func (d deck) print() {
	for i, card := range d{
		fmt.Println(i, card)
	}
}


func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// convert deck array to string, using JOIN.
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// write to file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666 )
}

func newDeckFromFile(filename string) deck {
	// read from file
	bs, err := ioutil.ReadFile(filename)

	// error handling
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// split a string by separator
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	/*shuffle cards inside the deck*/

	// get unixtime
	t := time.Now().UnixNano()

	// set a new random generator
	source := rand.NewSource(t)
	r := rand.New(source)

	// swap a card with 1 random card from the deck
	rand.Shuffle(len(d), func(i, j int) {
		j = r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	})
}