package main

import "math/rand"

// Deck holds the cards in the deck to be shuffled
type Deck []Card

// New creates a deck of cards to be used
func New() (deck Deck) {
	// Valid types include Two, Three, Four, Five, Six
	// Seven, Eight, Nine, Ten, Jack, Queen, King, Ace
	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	// Valid suits include Heart, Diamond, Club, Spade
	suits := []string{"\u2665", "\u2666", "\u2663", "\u2660"}

	// Loop over each type and suit appending to the deck
	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return
}

// Shuffle the deck
func Shuffle(deck Deck) Deck {
	for i := range deck {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the current card doesn't match the random
		// int we generated then we'll switch them out
		if i != r {
			deck[r], deck[i] = deck[i], deck[r]
		}
	}
	return deck
}

// Deal a specified amount of cards
func Deal(deck Deck, n int) ([]Card, Deck) {
	var cards []Card
	for i := 0; i < n; i++ {
		cards = append(cards, deck[i])
		deck = RemoveCard(deck, i, n)
	}
	return cards, deck
}
