package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Club, Rank: Ace})
	fmt.Println(Card{Suit: Joker})

	//Output:
	// Ace of Clubs
	// Joker

}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in deck")
		return
	}

}

func TestDefaultSort(t *testing.T) {
	input := []Card{
		{Suit: Club, Rank: Queen},
		{Suit: Spade, Rank: Six},
		{Suit: Club, Rank: Five},
		{Suit: Spade, Rank: Five},
		{Suit: Club, Rank: Six},
		{Suit: Spade, Rank: Queen},
	}
	want := []Card{
		{Suit: Spade, Rank: Five},
		{Suit: Spade, Rank: Six},
		{Suit: Spade, Rank: Queen},
		{Suit: Club, Rank: Five},
		{Suit: Club, Rank: Six},
		{Suit: Club, Rank: Queen},
	}
	got := DefaultSort(input)
	for i, v := range got {
		if got[i].Suit != want[i].Suit {
			t.Errorf("got: %v want: %v \n", v, want[i])
		} else if got[i].Rank != want[i].Rank {
			t.Errorf("got: %v want: %v \n", v, want[i])
		}
	}
}
