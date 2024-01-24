package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"
)

// Debug helps debugging the deck of cards
func Debug(deck Deck) {
	if os.Getenv("DEBUG") != "" {
		for i := range deck {
			fmt.Printf("Card #%d is a %s of %ss\n", i+1, deck[i].Type, deck[i].Suit)
		}
	}
}

func RemoveCard(deck Deck, index int, num int) Deck {
	slice := deck
	slice = slices.Delete(slice, index, index+num)
	return slice
}

func Hit(deck Deck, hand Hand) (Hand, Deck) {
	card, remainingDeck := Deal(deck, 1)
	newHand := append(hand, card...)
	return newHand, remainingDeck
}

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

func GetTotal(hand Hand) int {
	var total int
	var containsAce int
	for i := range hand {
		card := hand[i]
		if IsAceCard(&card) {
			containsAce += 1
		}
		total += CardValue(&card)
	}

	if containsAce > 0 && total > 21 {
		return total - 10
	}

	return total
}

func GetWinner(dealerTotal int, playerTotal int) string {
	if dealerTotal < 22 && dealerTotal > playerTotal {
		return "You lose :("
	} else if dealerTotal == playerTotal {
		return "Push"
	}
	return "You win!"
}

func ShowWelcome() {
	fmt.Println(strings.Repeat("#", 30))
	fmt.Println("####### \u2665 \u2666 Gojack \u2663 \u2660 #######")
	fmt.Println(strings.Repeat("#", 30) + "\n")
	// fmt.Println("Welcome!\n")
}
