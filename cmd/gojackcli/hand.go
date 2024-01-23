package main

type Hand []Card

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