package poker

import (
	"fmt"
	"math/rand"
	"sort"
)

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

type HandRank int

const (
	HighCard HandRank = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

var handRankToName = map[HandRank]string{
	HighCard:      "high card",
	OnePair:       "one pair",
	TwoPair:       "two pair",
	ThreeOfAKind:  "three of a kind",
	Straight:      "straight",
	Flush:         "flush",
	FullHouse:     "full house",
	FourOfAKind:   "four of a kind",
	StraightFlush: "straight flush",
	RoyalFlush:    "royal flush",
}

func (r HandRank) String() string {
	return handRankToName[r]
}

type Hand struct {
	rank  HandRank
	Cards []Card
}

func (h *Hand) sort() {
	sort.SliceStable(h.Cards, func(i, j int) bool {
		return h.Cards[i].Rank > h.Cards[j].Rank
	})
}

func (h *Hand) Rank() HandRank {
	if h.rank != 0 {
		return h.rank
	}
	h.sort()

	h.rank = HighCard

	if h.isFlush() {
		h.rank = Flush
		if h.isStraight() {
			h.rank = StraightFlush
			if h.Cards[0].Rank == Ace {
				h.rank = RoyalFlush
			}
		}
		return h.rank
	}

	if h.isStraight() {
		h.rank = Straight
		return h.rank
	}

	for i := 0; i < 4; i++ {
		if i < 2 {
			if h.Cards[i].Rank == h.Cards[i+1].Rank && h.Cards[i].Rank == h.Cards[i+2].Rank && h.Cards[i].Rank == h.Cards[i+3].Rank {
				h.rank = FourOfAKind
				return h.rank
			}
		}

		if i < 3 {
			if h.Cards[i].Rank == h.Cards[i+1].Rank && h.Cards[i].Rank == h.Cards[i+2].Rank {
				h.rank = ThreeOfAKind
				if i == 0 {
					if h.Cards[3].Rank == h.Cards[4].Rank {
						h.rank = FullHouse
					}
				}
				if i == 2 && h.rank == OnePair {
					h.rank = FullHouse
				}
				return h.rank
			}
		}

		if h.Cards[i].Rank == h.Cards[i+1].Rank {
			if h.rank == OnePair {
				return TwoPair
			}
			h.rank = OnePair
			i++
		}
	}

	return h.rank
}

func (h *Hand) isFlush() bool {
	s := h.Cards[0].Suit
	for _, c := range h.Cards[1:] {
		if c.Suit != s {
			return false
		}
	}
	return true
}

func (h *Hand) isStraight() bool {
	for i := 0; i < 4; i++ {
		// for ace to five straight
		if h.Cards[i].Rank == Ace && h.Cards[i+1].Rank == Five {
			continue
		}
		if h.Cards[i].Rank-1 != h.Cards[i+1].Rank {
			return false
		}
	}
	return true
}

type Card struct {
	Suit Suit
	Rank CardRank
}

func (c Card) String() string {
	return fmt.Sprintf("%d%s", c.Rank, c.Suit)
}

type Board struct {
	Cards []Card
}

type PersonalHand struct {
	Cards []Card
}

type Deck struct {
	Cards []Card
}

func (d *Deck) Draw() Card {
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c
}

func (d *Deck) Reset() {
	suits := []Suit{Spade, Club, Diamond, Heart}
	// 14 == ace
	ranks := []CardRank{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	d.Cards = make([]Card, 0, 52)
	for _, s := range suits {
		for _, r := range ranks {
			d.Cards = append(d.Cards, Card{Suit: s, Rank: r})
		}
	}

	rand.Shuffle(52, func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
