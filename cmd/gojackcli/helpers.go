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

func GetWinner(dealerTotal int, playerTotal int) string {
	if dealerTotal < 22 && dealerTotal > playerTotal {
		return "Dealer wins!"
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
