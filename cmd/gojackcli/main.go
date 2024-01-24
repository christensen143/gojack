package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	ShowWelcome()

	fmt.Printf("Dealer Hand: %v\n", dealerHand[0])
	fmt.Printf("Your Hand: %v (%v)\n", playerHand, GetTotal(playerHand))
	play := true

	if GetTotal(playerHand) == 21 && GetTotal(dealerHand) == 21 {
		fmt.Println("BlackJack Push!")
		return
	} else if GetTotal(playerHand) == 21 {
		fmt.Println("BlackJack! You win!")
		return
	} else if GetTotal(dealerHand) == 21 {
		fmt.Println("Dealer BlackJack! You lose!")
		return
	}

	for play {
		fmt.Print("Do you want to hit or stand?\n")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")

		if strings.ToLower(input) == "hit" {
			playerHandSlice, newDeckSlice = Hit(newDeckSlice, playerHandSlice)
			fmt.Printf("Your Hand: %v (%v)\n", playerHandSlice, GetTotal(playerHandSlice))
			if GetTotal(playerHandSlice) > 21 {
				fmt.Printf("Dealer Hand: %v (%v)\n", dealerHandSlice, GetTotal(dealerHandSlice))
				fmt.Println("Busted")
				return
			}
			// fmt.Printf("Dealer Total: %d\nYour Total: %d\n", GetTotal(dealerHandSlice), GetTotal(playerHandSlice))
		} else {
			play = false
			fmt.Printf("Dealer Hand: %v\n", dealerHandSlice)
		}
	}

	dealerTotal := GetTotal(dealerHandSlice)
	for dealerTotal < 17 {
		fmt.Println("Dealer Hits!")
		dealerHandSlice, newDeckSlice = Hit(newDeckSlice, dealerHandSlice)
		fmt.Printf("Dealer Hand: %v (%v)\n", dealerHandSlice, GetTotal(dealerHandSlice))
		dealerTotal = GetTotal(dealerHandSlice)
	}

	dealerTotal = GetTotal(dealerHandSlice)
	playerTotal := GetTotal(playerHandSlice)
	fmt.Printf("Dealer Total: %d\nYour Total: %d\n", dealerTotal, playerTotal)
	result := GetWinner(dealerTotal, playerTotal)
	fmt.Println(result)

}
