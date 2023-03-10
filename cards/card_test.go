package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 52 {
		t.Error("wrong number of cards in deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	want := Card{Rank: Ace, Suit: Spade}

	if cards[0] != want {
		t.Errorf("got: %v, want: %v", cards[0], want)
	}

}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}
