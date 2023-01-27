//go:generate stringer -type=Suit
//go:generate stringer -type=Rank

package deck

import (
	"fmt"
	"sort"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var Suits = [4]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	MinRank = Ace
	MaxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opt ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range Suits {
		for rank := MinRank; rank <= MaxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absCard(cards[i]) < absCard(cards[j])
	}
}

func absCard(c Card) int {
	return int(c.Suit)*int(MaxRank) + int(c.Rank)
}
