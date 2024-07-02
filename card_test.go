package deck

import "fmt"

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
