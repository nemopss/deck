package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Ten, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Ten of Diamonds
	// Jack of Clubs
	// Queen of Diamonds
	// King of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades as the first card, received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades as the first card, received:", cards[0])
	}
}
