package poker

import (
	"sort"
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

func NewHand(cards []Card) *Hand {
	h := &Hand{Cards: cards}
	h.Rank()
	return h
}

func (h *Hand) sortByCardRank() {
	sort.SliceStable(h.Cards, func(i, j int) bool {
		return h.Cards[i].Rank > h.Cards[j].Rank
	})
}

func (h *Hand) sortByHandRank() {
	switch h.rank {
	case StraightFlush, Straight:
		if h.Cards[0].Rank == Ace && h.Cards[1].Rank == Five {
			h.Cards = append(h.Cards[1:], h.Cards[0])
		}
	case FourOfAKind:
		if h.Cards[0].Rank != h.Cards[1].Rank {
			h.Cards = append(h.Cards[1:], h.Cards[0])
		}
	case FullHouse:
		if h.Cards[1].Rank != h.Cards[2].Rank {
			h.Cards = append(h.Cards[2:], h.Cards[:2]...)
		}
	case ThreeOfAKind:
		if h.Cards[0].Rank != h.Cards[1].Rank && h.Cards[1].Rank != h.Cards[2].Rank {
			h.Cards = append(h.Cards[2:], h.Cards[:2]...)
		}
		if h.Cards[0].Rank != h.Cards[1].Rank && h.Cards[3].Rank != h.Cards[4].Rank {
			h.Cards = append(h.Cards[1:4], h.Cards[0], h.Cards[4])
		}
	case TwoPair:
		if h.Cards[0].Rank != h.Cards[1].Rank {
			h.Cards = append(h.Cards[1:], h.Cards[0])
		}
		if h.Cards[1].Rank != h.Cards[2].Rank && h.Cards[2].Rank != h.Cards[3].Rank {
			h.Cards = append(h.Cards[:2], append(h.Cards[3:], h.Cards[2])...)
		}
	case OnePair:
		if h.Cards[1].Rank == h.Cards[2].Rank {
			h.Cards = append(h.Cards[1:3], append([]Card{h.Cards[0]}, h.Cards[3:]...)...)
		}
		if h.Cards[2].Rank == h.Cards[3].Rank {
			h.Cards = append(append(h.Cards[2:4], h.Cards[:2]...), h.Cards[4])
		}
		if h.Cards[3].Rank == h.Cards[4].Rank {
			h.Cards = append(h.Cards[3:], h.Cards[:3]...)
		}
	}
}

func (h *Hand) Rank() HandRank {
	if h.rank != 0 {
		return h.rank
	}
	h.sortByCardRank()
	defer h.sortByHandRank()

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

type Result string

const (
	Win  Result = "win"
	Lose Result = "lose"
	Draw Result = "draw"
)

func (h *Hand) Compare(rival *Hand) Result {
	if h.Rank() > rival.Rank() {
		return Win
	}
	if h.Rank() < rival.Rank() {
		return Lose
	}

	if h.Cards[0].Rank > rival.Cards[0].Rank {
		return Win
	}
	if h.Cards[0].Rank < rival.Cards[0].Rank {
		return Lose
	}

	switch h.Rank() {
	case FourOfAKind:
		if h.Cards[4].Rank > rival.Cards[4].Rank {
			return Win
		}
		if h.Cards[4].Rank < rival.Cards[4].Rank {
			return Lose
		}
		return Draw
	case FullHouse:
		if h.Cards[3].Rank > rival.Cards[3].Rank {
			return Win
		}
		if h.Cards[3].Rank < rival.Cards[3].Rank {
			return Lose
		}
		return Draw
	case Flush:
		for i := 1; i < 5; i++ {
			if h.Cards[i].Rank > rival.Cards[i].Rank {
				return Win
			}
			if h.Cards[i].Rank < rival.Cards[i].Rank {
				return Lose
			}
		}
		return Draw
	case ThreeOfAKind:
		for i := 3; i < 5; i++ {
			if h.Cards[i].Rank > rival.Cards[i].Rank {
				return Win
			}
			if h.Cards[i].Rank < rival.Cards[i].Rank {
				return Lose
			}
		}
		return Draw
	case TwoPair:
		for i := 2; i < 5; i += 2 {
			if h.Cards[i].Rank > rival.Cards[i].Rank {
				return Win
			}
			if h.Cards[i].Rank < rival.Cards[i].Rank {
				return Lose
			}
		}
		return Draw
	case OnePair:
		for i := 2; i < 5; i++ {
			if h.Cards[i].Rank > rival.Cards[i].Rank {
				return Win
			}
			if h.Cards[i].Rank < rival.Cards[i].Rank {
				return Lose
			}
		}
		return Draw
	case HighCard:
		for i := 1; i < 5; i++ {
			if h.Cards[i].Rank > rival.Cards[i].Rank {
				return Win
			}
			if h.Cards[i].Rank < rival.Cards[i].Rank {
				return Lose
			}
		}
		return Draw
	default:
		return Draw
	}
}

type Board struct {
	Cards []Card
}

type PersonalHand struct {
	Cards []Card
}
