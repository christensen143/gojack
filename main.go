package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"
)

// Card holds the card suits and types in the deck
type Card struct {
	Type string
	Suit string
}

type Hand []Card

// Deck holds the cards in the deck to be shuffled
type Deck []Card

func CardValue(c Card) int {
	switch c.Type {
	case "Two":
		return 2
	case "Three":
		return 3
	case "Four":
		return 4
	case "Five":
		return 5
	case "Six":
		return 6
	case "Seven":
		return 7
	case "Eight":
		return 8
	case "Nine":
		return 9
	case "Ten", "Jack", "Queen", "King":
		return 10
	case "Ace":
		return 11
	default:
		return 0
	}
}

func IsAceCard(c Card) bool {
	return c.Type == "Ace"
}

func GetTotal(h Hand) int {
	var total int
	var containsAce int
	for i := 0; i < len(h); i++ {
		card := h[i]
		if IsAceCard(card) {
			containsAce += 1
		}
		total += CardValue(card)
	}

	if containsAce > 0 && total > 21 {
		return total - 10
	}

	return total
}

// New creates a deck of cards to be used
func New() (deck Deck) {
	// Valid types include Two, Three, Four, Five, Six
	// Seven, Eight, Nine, Ten, Jack, Queen, King, Ace
	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	// Valid suits include Heart, Diamond, Club, Spade
	suits := []string{"Heart", "Diamond", "Club", "Spade"}

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
func Shuffle(d Deck) Deck {
	for i := 1; i < len(d); i++ {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the current card doesn't match the random
		// int we generated then we'll switch them out
		if i != r {
			d[r], d[i] = d[i], d[r]
		}
	}
	return d
}

// Deal a specified amount of cards
func Deal(d Deck, n int) ([]Card, Deck) {
	var cards []Card
	deck := d
	for i := 0; i < n; i++ {
		cards = append(cards, d[i])
		deck = RemoveCard(d, i, n)
	}
	return cards, deck
}

func GetHands(cards []Card) ([]Card, []Card) {
	var dealerHand []Card
	var playerHand []Card
	for c := range cards {
		if c == 0 || c == 2 {
			dealerHand = append(dealerHand, cards[c])
		} else {
			playerHand = append(playerHand, cards[c])
		}
	}
	return dealerHand, playerHand
}

// Debug helps debugging the deck of cards
func Debug(d Deck) {
	if os.Getenv("DEBUG") != "" {
		for i := 0; i < len(d); i++ {
			fmt.Printf("Card #%d is a %s of %ss\n", i+1, d[i].Type, d[i].Suit)
		}
	}
}

func RemoveCard(s Deck, index int, num int) Deck {
	slice := s
	slice = slices.Delete(slice, index, index+num)
	return slice
}

func Hit(d Deck, h Hand) (Hand, Deck) {
	card, deck := Deal(d, 1)
	newHand := append(h, card...)
	return newHand, deck
}

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

func main() {
	deck := New()
	Debug(deck)
	Shuffle(deck)
	Debug(deck)
	cards, newDeck := Deal(deck, 4)
	dealerHand, playerHand := GetHands(cards)
	playerHandSlice := playerHand
	dealerHandSlice := dealerHand
	newDeckSlice := newDeck
	fmt.Printf("Dealer Hand: %v\n", dealerHand[0])
	fmt.Printf("Player Hand: %v\n", playerHand)

	play := true

	for play {
		fmt.Print("Do you want to hit or stay?\n")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")

		if strings.ToLower(input) == "hit" {
			playerHandSlice, newDeckSlice = Hit(newDeckSlice, playerHandSlice)
			fmt.Printf("Player Hand: %v\n", playerHandSlice)
			if GetTotal(playerHandSlice) > 21 {
				fmt.Printf("Dealer Hand: %v\n", dealerHandSlice)
				fmt.Println("Busted")
				return
			}
			// fmt.Printf("Dealer Total: %d\nPlayer Total: %d\n", GetTotal(dealerHandSlice), GetTotal(playerHandSlice))
		} else {
			play = false
			fmt.Printf("Dealer Hand: %v\n", dealerHandSlice)
		}
	}

	dealerTotal := GetTotal(dealerHandSlice)
	for dealerTotal < 17 {
		dealerHandSlice, newDeckSlice = Hit(newDeckSlice, dealerHandSlice)
		fmt.Printf("Dealer Hand: %v\n", dealerHandSlice)
		dealerTotal = GetTotal(dealerHandSlice)
	}

	dealerTotal = GetTotal(dealerHandSlice)
	playerTotal := GetTotal(playerHandSlice)
	fmt.Printf("Dealer Total: %d\nPlayer Total: %d\n", dealerTotal, playerTotal)

}
