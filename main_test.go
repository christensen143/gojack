package main

import "testing"

func TestCardValue(t *testing.T) {
  testCases := []struct {
    card Card
    expected int 
  }{
    {Card{"Two", "Spades"}, 2},
    {Card{"Ten", "Hearts"}, 10},
    {Card{"Ace", "Clubs"}, 11},
  }

  for _, tc := range testCases {
    actual := CardValue(tc.card)
    if actual != tc.expected {
      t.Errorf("CardValue(%v) expected %d, actual %d", tc.card, tc.expected, actual) 
    }
  }
}

func TestGetTotal(t *testing.T) {
  hand := Hand{
    Card{"Ace", "Spades"},
    Card{"King", "Hearts"},
  }
  
  expected := 21
  actual := GetTotal(hand)
  
  if actual != expected {
    t.Errorf("GetTotal(%v) expected %d, actual %d", hand, expected, actual)
  }
}

func TestGetHands(t *testing.T) {
  cards := []Card{
    {"Ace", "Spades"}, 
    {"King", "Hearts"},
    {"Queen", "Clubs"},
    {"Jack", "Diamonds"},
  }
  
  expectedDealer := []Card{{"Ace", "Spades"}, {"Queen", "Clubs"}}
  expectedPlayer := []Card{{"King", "Hearts"}, {"Jack", "Diamonds"}}
  
  dealer, player := GetHands(cards)
  
  if len(dealer) != len(expectedDealer) || len(player) != len(expectedPlayer) {
    t.Errorf("GetHands(%v) returned incorrect hand sizes", cards)
  }
  
  for i := range dealer {
    if dealer[i] != expectedDealer[i] {
      t.Errorf("GetHands(%v) returned incorrect dealer hand", cards) 
    }
  }
  
  for i := range player {
    if player[i] != expectedPlayer[i] {
      t.Errorf("GetHands(%v) returned incorrect player hand", cards)
    }
  }
}
