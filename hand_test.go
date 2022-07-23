package poker

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHand_sortByCardRank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hand *Hand
		want *Hand
	}{
		{
			name: "normal",
			hand: &Hand{
				Cards: []Card{{Rank: 4}, {Rank: 9}, {Rank: 2}, {Rank: 8}, {Rank: 5}},
			},
			want: &Hand{
				Cards: []Card{{Rank: 9}, {Rank: 8}, {Rank: 5}, {Rank: 4}, {Rank: 2}},
			},
		},
		{
			name: "pair",
			hand: &Hand{
				Cards: []Card{{Rank: 4}, {Rank: 2}, {Rank: 6}, {Rank: 4}, {Rank: 9}},
			},
			want: &Hand{
				Cards: []Card{{Rank: 9}, {Rank: 6}, {Rank: 4}, {Rank: 4}, {Rank: 2}},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.hand.sortByCardRank()
			if diff := cmp.Diff(tt.hand.Cards, tt.want.Cards); diff != "" {
				t.Fatalf("want and got are different(-got +want): %s", diff)
			}
		})
	}
}

func TestHand_sortByHandRank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hand *Hand
		want *Hand
	}{
		{
			name: "straight flush",
			hand: &Hand{
				rank:  StraightFlush,
				Cards: []Card{{Rank: Ace, Suit: Spade}, {Rank: King, Suit: Spade}, {Rank: Queen, Suit: Spade}, {Rank: Jack, Suit: Spade}, {Rank: Ten, Suit: Spade}},
			},
			want: &Hand{
				rank:  StraightFlush,
				Cards: []Card{{Rank: Ace, Suit: Spade}, {Rank: King, Suit: Spade}, {Rank: Queen, Suit: Spade}, {Rank: Jack, Suit: Spade}, {Rank: Ten, Suit: Spade}},
			},
		},
		{
			name: "straight flush 5~1",
			hand: &Hand{
				rank:  StraightFlush,
				Cards: []Card{{Rank: Ace, Suit: Spade}, {Rank: Five, Suit: Spade}, {Rank: Four, Suit: Spade}, {Rank: Three, Suit: Spade}, {Rank: Deuce, Suit: Spade}},
			},
			want: &Hand{
				rank:  StraightFlush,
				Cards: []Card{{Rank: Five, Suit: Spade}, {Rank: Four, Suit: Spade}, {Rank: Three, Suit: Spade}, {Rank: Deuce, Suit: Spade}, {Rank: Ace, Suit: Spade}},
			},
		},
		{
			name: "straight",
			hand: &Hand{
				rank:  Straight,
				Cards: []Card{{Rank: Ace}, {Rank: King}, {Rank: Queen}, {Rank: Jack}, {Rank: Ten}},
			},
			want: &Hand{
				rank:  Straight,
				Cards: []Card{{Rank: Ace}, {Rank: King}, {Rank: Queen}, {Rank: Jack}, {Rank: Ten}},
			},
		},
		{
			name: "straight 5~1",
			hand: &Hand{
				rank:  Straight,
				Cards: []Card{{Rank: Ace}, {Rank: Five}, {Rank: Four}, {Rank: Three}, {Rank: Deuce}},
			},
			want: &Hand{
				rank:  Straight,
				Cards: []Card{{Rank: Five}, {Rank: Four}, {Rank: Three}, {Rank: Deuce}, {Rank: Ace}},
			},
		},
		{
			name: "four of a kind",
			hand: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: Five}, {Rank: Five}, {Rank: Five}, {Rank: Five}, {Rank: Deuce}},
			},
			want: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: Five}, {Rank: Five}, {Rank: Five}, {Rank: Five}, {Rank: Deuce}},
			},
		},
		{
			name: "four of a kind 2",
			hand: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: Ten}, {Rank: Five}, {Rank: Five}, {Rank: Five}, {Rank: Five}},
			},
			want: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: Five}, {Rank: Five}, {Rank: Five}, {Rank: Five}, {Rank: Ten}},
			},
		},
		{
			name: "full house",
			hand: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: Five}, {Rank: Five}},
			},
			want: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: Five}, {Rank: Five}},
			},
		},
		{
			name: "full house 2",
			hand: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Five}, {Rank: Five}, {Rank: Five}},
			},
			want: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Five}, {Rank: Five}, {Rank: Five}, {Rank: Ten}, {Rank: Ten}},
			},
		},
		{
			name: "three of a kind",
			hand: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: Six}, {Rank: Five}},
			},
			want: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: Six}, {Rank: Five}},
			},
		},
		{
			name: "three of a kind 2",
			hand: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Jack}, {Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: Five}},
			},
			want: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: Jack}, {Rank: Five}},
			},
		},
		{
			name: "three of a kind 3",
			hand: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: King}, {Rank: Jack}, {Rank: Ten}, {Rank: Ten}, {Rank: Ten}},
			},
			want: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: King}, {Rank: Jack}},
			},
		},
		{
			name: "two pair",
			hand: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Jack}, {Rank: Jack}, {Rank: Eight}, {Rank: Eight}, {Rank: Deuce}},
			},
			want: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Jack}, {Rank: Jack}, {Rank: Eight}, {Rank: Eight}, {Rank: Deuce}},
			},
		},
		{
			name: "two pair 2",
			hand: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: King}, {Rank: Jack}, {Rank: Jack}, {Rank: Eight}, {Rank: Eight}},
			},
			want: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Jack}, {Rank: Jack}, {Rank: Eight}, {Rank: Eight}, {Rank: King}},
			},
		},
		{
			name: "two pair 3",
			hand: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Jack}, {Rank: Jack}, {Rank: Nine}, {Rank: Eight}, {Rank: Eight}},
			},
			want: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Jack}, {Rank: Jack}, {Rank: Eight}, {Rank: Eight}, {Rank: Nine}},
			},
		},
		{
			name: "one pair",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Jack}, {Rank: Five}, {Rank: Three}},
			},
			want: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Jack}, {Rank: Five}, {Rank: Three}},
			},
		},
		{
			name: "one pair 2",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Jack}, {Rank: Jack}, {Rank: Five}, {Rank: Three}},
			},
			want: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Jack}, {Rank: Jack}, {Rank: Ace}, {Rank: Five}, {Rank: Three}},
			},
		},
		{
			name: "one pair 3",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Jack}, {Rank: Five}, {Rank: Five}, {Rank: Three}},
			},
			want: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Five}, {Rank: Five}, {Rank: Ace}, {Rank: Jack}, {Rank: Three}},
			},
		},
		{
			name: "one pair 4",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Jack}, {Rank: Five}, {Rank: Three}, {Rank: Three}},
			},
			want: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Three}, {Rank: Three}, {Rank: Ace}, {Rank: Jack}, {Rank: Five}},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.hand.sortByHandRank()
			if diff := cmp.Diff(tt.hand.Cards, tt.want.Cards); diff != "" {
				t.Fatalf("sort is wrong(-got +want):\n%s", diff)
			}
		})
	}
}

func TestHand_isFlush(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hand *Hand
		want bool
	}{
		{
			name: "flush",
			hand: &Hand{
				Cards: []Card{{Suit: Spade}, {Suit: Spade}, {Suit: Spade}, {Suit: Spade}, {Suit: Spade}},
			},
			want: true,
		},
		{
			name: "not flush",
			hand: &Hand{
				Cards: []Card{{Suit: Spade}, {Suit: Spade}, {Suit: Spade}, {Suit: Spade}, {Suit: Heart}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.hand.isFlush() != tt.want {
				t.Fatalf("isFlush is wrong")
			}
		})
	}
}

func TestHand_isStraight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hand *Hand
		want bool
	}{
		{
			name: "straight1~5",
			hand: &Hand{
				Cards: []Card{{Rank: Ace}, {Rank: Five}, {Rank: Four}, {Rank: Three}, {Rank: Deuce}},
			},
			want: true,
		},
		{
			name: "straight1~10",
			hand: &Hand{
				Cards: []Card{{Rank: Ace}, {Rank: King}, {Rank: Queen}, {Rank: Jack}, {Rank: Ten}},
			},
			want: true,
		},
		{
			name: "straight5~9",
			hand: &Hand{
				Cards: []Card{{Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}, {Rank: Five}},
			},
			want: true,
		},
		{
			name: "not straight",
			hand: &Hand{
				Cards: []Card{{Rank: Eight}, {Rank: Eight}, {Rank: Three}, {Rank: Deuce}, {Rank: Deuce}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.hand.isStraight() != tt.want {
				t.Fatalf("isStraight is wrong")
			}
		})
	}
}

func TestHand_Rank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hand *Hand
		want HandRank
	}{
		{
			name: "royal flush",
			hand: &Hand{
				Cards: []Card{
					{Suit: Spade, Rank: Ace},
					{Suit: Spade, Rank: King},
					{Suit: Spade, Rank: Queen},
					{Suit: Spade, Rank: Jack},
					{Suit: Spade, Rank: Ten},
				},
			},
			want: RoyalFlush,
		},
		{
			name: "straight flush",
			hand: &Hand{
				Cards: []Card{
					{Suit: Club, Rank: Seven},
					{Suit: Club, Rank: Six},
					{Suit: Club, Rank: Five},
					{Suit: Club, Rank: Four},
					{Suit: Club, Rank: Three},
				},
			},
			want: StraightFlush,
		},
		{
			name: "four of a kind",
			hand: &Hand{
				Cards: []Card{
					{Suit: Club, Rank: Ace},
					{Suit: Spade, Rank: Five},
					{Suit: Club, Rank: Five},
					{Suit: Diamond, Rank: Five},
					{Suit: Heart, Rank: Five},
				},
			},
			want: FourOfAKind,
		},
		{
			name: "full house",
			hand: &Hand{
				Cards: []Card{
					{Suit: Spade, Rank: Jack},
					{Suit: Club, Rank: Jack},
					{Suit: Heart, Rank: Jack},
					{Suit: Diamond, Rank: Four},
					{Suit: Diamond, Rank: Four},
				},
			},
			want: FullHouse,
		},
		{
			name: "flush",
			hand: &Hand{
				Cards: []Card{
					{Suit: Heart, Rank: Ace},
					{Suit: Heart, Rank: Jack},
					{Suit: Heart, Rank: Nine},
					{Suit: Heart, Rank: Three},
					{Suit: Heart, Rank: Deuce},
				},
			},
			want: Flush,
		},
		{
			name: "straight",
			hand: &Hand{
				Cards: []Card{
					{Suit: Diamond, Rank: Jack},
					{Suit: Diamond, Rank: Ten},
					{Suit: Spade, Rank: Nine},
					{Suit: Heart, Rank: Eight},
					{Suit: Heart, Rank: Seven},
				},
			},
			want: Straight,
		},
		{
			name: "three of a kind",
			hand: &Hand{
				Cards: []Card{
					{Suit: Heart, Rank: Queen},
					{Suit: Spade, Rank: Ten},
					{Suit: Spade, Rank: Seven},
					{Suit: Diamond, Rank: Seven},
					{Suit: Heart, Rank: Seven},
				},
			},
			want: ThreeOfAKind,
		},
		{
			name: "two pair",
			hand: &Hand{
				Cards: []Card{
					{Suit: Spade, Rank: Ten},
					{Suit: Diamond, Rank: Ten},
					{Suit: Diamond, Rank: Six},
					{Suit: Spade, Rank: Deuce},
					{Suit: Heart, Rank: Deuce},
				},
			},
			want: TwoPair,
		},
		{
			name: "one pair",
			hand: &Hand{
				Cards: []Card{
					{Suit: Heart, Rank: Ace},
					{Suit: Diamond, Rank: Eight},
					{Suit: Heart, Rank: Eight},
					{Suit: Spade, Rank: Seven},
					{Suit: Club, Rank: Three},
				},
			},
			want: OnePair,
		},
		{
			name: "high card",
			hand: &Hand{
				Cards: []Card{
					{Suit: Club, Rank: King},
					{Suit: Spade, Rank: Ten},
					{Suit: Spade, Rank: Six},
					{Suit: Heart, Rank: Four},
					{Suit: Club, Rank: Deuce},
				},
			},
			want: HighCard,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.hand.Rank() != tt.want {
				t.Fatalf("want is %s, but got %s", tt.want, tt.hand.Rank())
			}
		})
	}
}

func TestHand_Compare_HandRank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		handRank HandRank
		rivals   []HandRank
		want     Result
	}{
		{
			name:     "royal flush win",
			handRank: RoyalFlush,
			rivals:   []HandRank{StraightFlush, FourOfAKind, FullHouse, Flush, Straight, ThreeOfAKind, TwoPair, OnePair, HighCard},
			want:     Win,
		},
		{
			name:     "straight flush win",
			handRank: StraightFlush,
			rivals:   []HandRank{FourOfAKind, FullHouse, Flush, Straight, ThreeOfAKind, TwoPair, OnePair, HighCard},
			want:     Win,
		},
		{
			name:     "four of a kind win",
			handRank: FourOfAKind,
			rivals:   []HandRank{FullHouse, Flush, Straight, ThreeOfAKind, TwoPair, OnePair, HighCard},
			want:     Win,
		},
		{
			name:     "full house win",
			handRank: FullHouse,
			rivals:   []HandRank{Flush, Straight, ThreeOfAKind, TwoPair, OnePair, HighCard},
			want:     Win,
		},
		{
			name:     "flush win",
			handRank: Flush,
			rivals:   []HandRank{Straight, ThreeOfAKind, TwoPair, OnePair, HighCard},
			want:     Win,
		},
		{
			name:     "straight win",
			handRank: Straight,
			rivals:   []HandRank{ThreeOfAKind, TwoPair, OnePair, HighCard},
			want:     Win,
		},
		{
			name:     "three of a kind win",
			handRank: ThreeOfAKind,
			rivals:   []HandRank{TwoPair, OnePair, HighCard},
			want:     Win,
		},
		{
			name:     "two pair win",
			handRank: TwoPair,
			rivals:   []HandRank{OnePair, HighCard},
			want:     Win,
		},
		{
			name:     "one pair win",
			handRank: OnePair,
			rivals:   []HandRank{HighCard},
			want:     Win,
		},
		{
			name:     "high card lose",
			handRank: HighCard,
			rivals:   []HandRank{RoyalFlush, StraightFlush, FourOfAKind, FullHouse, Flush, Straight, ThreeOfAKind, TwoPair, OnePair},
			want:     Lose,
		},
		{
			name:     "one pair lose",
			handRank: OnePair,
			rivals:   []HandRank{RoyalFlush, StraightFlush, FourOfAKind, FullHouse, Flush, Straight, ThreeOfAKind, TwoPair},
			want:     Lose,
		},
		{
			name:     "two pair lose",
			handRank: TwoPair,
			rivals:   []HandRank{RoyalFlush, StraightFlush, FourOfAKind, FullHouse, Flush, Straight, ThreeOfAKind},
			want:     Lose,
		},
		{
			name:     "three of a kind lose",
			handRank: ThreeOfAKind,
			rivals:   []HandRank{RoyalFlush, StraightFlush, FourOfAKind, FullHouse, Flush, Straight},
			want:     Lose,
		},
		{
			name:     "straight lose",
			handRank: Straight,
			rivals:   []HandRank{RoyalFlush, StraightFlush, FourOfAKind, FullHouse, Flush},
			want:     Lose,
		},
		{
			name:     "flush lose",
			handRank: Flush,
			rivals:   []HandRank{RoyalFlush, StraightFlush, FourOfAKind, FullHouse},
			want:     Lose,
		},
		{
			name:     "full house lose",
			handRank: FullHouse,
			rivals:   []HandRank{RoyalFlush, StraightFlush, FourOfAKind},
			want:     Lose,
		},
		{
			name:     "four of a kind lose",
			handRank: FourOfAKind,
			rivals:   []HandRank{RoyalFlush, StraightFlush},
			want:     Lose,
		},
		{
			name:     "straight flush lose",
			handRank: StraightFlush,
			rivals:   []HandRank{RoyalFlush},
			want:     Lose,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := &Hand{rank: tt.handRank}
			for _, r := range tt.rivals {
				if result := h.Compare(&Hand{rank: r}); result != tt.want {
					t.Fatalf("%s %s to %s", h.rank, result, r)
				}
			}
		})
	}
}

func TestHand_Compare_CardRank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		hand  *Hand
		rival *Hand
		want  Result
	}{
		{
			name: "high card win",
			hand: &Hand{
				rank:  HighCard,
				Cards: []Card{{Rank: Ace}, {Rank: Jack}, {Rank: Eight}, {Rank: Seven}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  HighCard,
				Cards: []Card{{Rank: Jack}, {Rank: Ten}, {Rank: Eight}, {Rank: Seven}, {Rank: Five}},
			},
			want: Win,
		},
		{
			name: "high card win 2",
			hand: &Hand{
				rank:  HighCard,
				Cards: []Card{{Rank: Ace}, {Rank: Jack}, {Rank: Eight}, {Rank: Seven}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  HighCard,
				Cards: []Card{{Rank: Ace}, {Rank: Ten}, {Rank: Eight}, {Rank: Seven}, {Rank: Five}},
			},
			want: Win,
		},
		{
			name: "high card draw",
			hand: &Hand{
				rank:  HighCard,
				Cards: []Card{{Rank: Ace}, {Rank: Jack}, {Rank: Eight}, {Rank: Seven}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  HighCard,
				Cards: []Card{{Rank: Ace}, {Rank: Jack}, {Rank: Eight}, {Rank: Seven}, {Rank: Five}},
			},
			want: Draw,
		},
		{
			name: "one pair win",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Nine}, {Rank: Five}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Nine}, {Rank: Nine}, {Rank: Eight}, {Rank: Five}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "one pair win 2",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Nine}, {Rank: Five}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Eight}, {Rank: Five}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "one pair win 3",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Nine}, {Rank: Five}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Nine}, {Rank: Four}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "one pair win 4",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Nine}, {Rank: Five}, {Rank: Three}},
			},
			rival: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Nine}, {Rank: Five}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "one pair draw",
			hand: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Eight}, {Rank: Five}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  OnePair,
				Cards: []Card{{Rank: Ace}, {Rank: Ace}, {Rank: Eight}, {Rank: Five}, {Rank: Deuce}},
			},
			want: Draw,
		},
		{
			name: "two pair win",
			hand: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Queen}, {Rank: Queen}, {Rank: Five}, {Rank: Five}, {Rank: Ten}},
			},
			rival: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Five}, {Rank: Five}, {Rank: Queen}},
			},
			want: Win,
		},
		{
			name: "two pair win 2",
			hand: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Queen}, {Rank: Queen}, {Rank: Ten}, {Rank: Ten}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Queen}, {Rank: Queen}, {Rank: Five}, {Rank: Five}, {Rank: Ten}},
			},
			want: Win,
		},
		{
			name: "two pair win 3",
			hand: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Queen}, {Rank: Queen}, {Rank: Ten}, {Rank: Ten}, {Rank: Jack}},
			},
			rival: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Queen}, {Rank: Queen}, {Rank: Ten}, {Rank: Ten}, {Rank: Nine}},
			},
			want: Win,
		},
		{
			name: "two pair draw",
			hand: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Queen}, {Rank: Queen}, {Rank: Ten}, {Rank: Ten}, {Rank: Jack}},
			},
			rival: &Hand{
				rank:  TwoPair,
				Cards: []Card{{Rank: Queen}, {Rank: Queen}, {Rank: Ten}, {Rank: Ten}, {Rank: Jack}},
			},
			want: Draw,
		},
		{
			name: "three of a kind win",
			hand: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Eight}, {Rank: Eight}, {Rank: Eight}, {Rank: Jack}, {Rank: Four}},
			},
			rival: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Four}, {Rank: Four}, {Rank: Four}, {Rank: Jack}, {Rank: Eight}},
			},
			want: Win,
		},
		{
			name: "three of a kind win 2",
			hand: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Eight}, {Rank: Eight}, {Rank: Eight}, {Rank: Jack}, {Rank: Four}},
			},
			rival: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Eight}, {Rank: Eight}, {Rank: Eight}, {Rank: Ten}, {Rank: Four}},
			},
			want: Win,
		},
		{
			name: "three of a kind win 3",
			hand: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Eight}, {Rank: Eight}, {Rank: Eight}, {Rank: Ten}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Eight}, {Rank: Eight}, {Rank: Eight}, {Rank: Ten}, {Rank: Four}},
			},
			want: Win,
		},
		{
			name: "three of a kind draw",
			hand: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Eight}, {Rank: Eight}, {Rank: Eight}, {Rank: Ten}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  ThreeOfAKind,
				Cards: []Card{{Rank: Eight}, {Rank: Eight}, {Rank: Eight}, {Rank: Ten}, {Rank: Five}},
			},
			want: Draw,
		},
		{
			name: "straight win",
			hand: &Hand{
				rank:  Straight,
				Cards: []Card{{Rank: Ten}, {Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}},
			},
			rival: &Hand{
				rank:  Straight,
				Cards: []Card{{Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}, {Rank: Five}},
			},
			want: Win,
		},
		{
			name: "straight draw",
			hand: &Hand{
				rank:  Straight,
				Cards: []Card{{Rank: Ten}, {Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}},
			},
			rival: &Hand{
				rank:  Straight,
				Cards: []Card{{Rank: Ten}, {Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}},
			},
			want: Draw,
		},
		{
			name: "flush win",
			hand: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Nine}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Ten}, {Rank: Nine}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "flush win 2",
			hand: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Nine}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "flush win 3",
			hand: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Nine}, {Rank: Four}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "flush win 4",
			hand: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Eight}, {Rank: Six}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "flush win 5",
			hand: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Eight}, {Rank: Four}, {Rank: Three}},
			},
			rival: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			want: Win,
		},
		{
			name: "flush draw",
			hand: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			rival: &Hand{
				rank:  Flush,
				Cards: []Card{{Rank: Queen}, {Rank: Ten}, {Rank: Eight}, {Rank: Four}, {Rank: Deuce}},
			},
			want: Draw,
		},
		{
			name: "full house win",
			hand: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Nine}, {Rank: Nine}, {Rank: Nine}, {Rank: Six}, {Rank: Six}},
			},
			rival: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Six}, {Rank: Six}, {Rank: Six}, {Rank: Nine}, {Rank: Nine}},
			},
			want: Win,
		},
		{
			name: "full house win 2",
			hand: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Nine}, {Rank: Nine}, {Rank: Nine}, {Rank: Six}, {Rank: Six}},
			},
			rival: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Nine}, {Rank: Nine}, {Rank: Nine}, {Rank: Five}, {Rank: Five}},
			},
			want: Win,
		},
		{
			name: "full house draw",
			hand: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Nine}, {Rank: Nine}, {Rank: Nine}, {Rank: Six}, {Rank: Six}},
			},
			rival: &Hand{
				rank:  FullHouse,
				Cards: []Card{{Rank: Nine}, {Rank: Nine}, {Rank: Nine}, {Rank: Six}, {Rank: Six}},
			},
			want: Draw,
		},
		{
			name: "four of a kind win",
			hand: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: King}, {Rank: King}, {Rank: King}, {Rank: King}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: Ten}, {Rank: Five}},
			},
			want: Win,
		},
		{
			name: "four of a kind win 2",
			hand: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: King}, {Rank: King}, {Rank: King}, {Rank: King}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: King}, {Rank: King}, {Rank: King}, {Rank: King}, {Rank: Four}},
			},
			want: Win,
		},
		{
			name: "four of a kind draw",
			hand: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: King}, {Rank: King}, {Rank: King}, {Rank: King}, {Rank: Five}},
			},
			rival: &Hand{
				rank:  FourOfAKind,
				Cards: []Card{{Rank: King}, {Rank: King}, {Rank: King}, {Rank: King}, {Rank: Five}},
			},
			want: Draw,
		},
		{
			name: "straight flush win",
			hand: &Hand{
				rank:  StraightFlush,
				Cards: []Card{{Rank: Ten}, {Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}},
			},
			rival: &Hand{
				rank:  StraightFlush,
				Cards: []Card{{Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}, {Rank: Five}},
			},
			want: Win,
		},
		{
			name: "straight flush draw",
			hand: &Hand{
				rank:  StraightFlush,
				Cards: []Card{{Rank: Ten}, {Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}},
			},
			rival: &Hand{
				rank:  StraightFlush,
				Cards: []Card{{Rank: Ten}, {Rank: Nine}, {Rank: Eight}, {Rank: Seven}, {Rank: Six}},
			},
			want: Draw,
		},
		{
			name: "royal flush draw",
			hand: &Hand{
				rank:  RoyalFlush,
				Cards: []Card{{Rank: Ace}, {Rank: King}, {Rank: Queen}, {Rank: Jack}, {Rank: Ten}},
			},
			rival: &Hand{
				rank:  RoyalFlush,
				Cards: []Card{{Rank: Ace}, {Rank: King}, {Rank: Queen}, {Rank: Jack}, {Rank: Ten}},
			},
			want: Draw,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if r := tt.hand.Compare(tt.rival); r != tt.want {
				t.Fatalf("hand %s to rival", r)
			}
			if tt.want == Win {
				if r := tt.rival.Compare(tt.hand); r != Lose {
					t.Fatalf("rival %s to hand", r)
				}
			}
			if tt.want == Draw {
				if r := tt.rival.Compare(tt.hand); r != Draw {
					t.Fatalf("rival %s to hand", r)
				}
			}
		})
	}
}
