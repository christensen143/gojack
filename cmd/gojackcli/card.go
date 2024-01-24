package main

// Card holds the card suits and types in the deck
type Card struct {
	Type string
	Suit string
}

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
