package deck

import (
	"fmt"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Club, Rank: Ace})
	fmt.Println(Card{Suit: Joker})

	//Output:
	// Ace of Clubs
	// Joker

}
