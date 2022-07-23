package poker

import "fmt"

type Suit string

const (
	Spade   Suit = "s"
	Club    Suit = "c"
	Diamond Suit = "d"
	Heart   Suit = "h"
)

type CardRank int

const (
	Ace CardRank = 14 - iota
	King
	Queen
	Jack
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Deuce
)

type Card struct {
	Suit Suit
	Rank CardRank
}

func (c Card) String() string {
	if c.Rank == Ace {
		return fmt.Sprintf("%d%s", 1, c.Suit)
	}
	return fmt.Sprintf("%d%s", c.Rank, c.Suit)
}
