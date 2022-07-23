package poker

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestDeck_Draw(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		d        *Deck
		wantCard Card
		wantDeck *Deck
	}{
		{
			name: "normal",
			d: &Deck{
				Cards: []Card{
					{Suit: Spade, Rank: 2},
					{Suit: Heart, Rank: 4},
				},
			},
			wantCard: Card{Suit: Spade, Rank: 2},
			wantDeck: &Deck{
				Cards: []Card{{Suit: Heart, Rank: 4}},
			},
		},
		{
			name: "last",
			d: &Deck{
				Cards: []Card{{Suit: Spade, Rank: 2}},
			},
			wantCard: Card{Suit: Spade, Rank: 2},
			wantDeck: &Deck{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := tt.d.Draw()
			if !cmp.Equal(c, tt.wantCard) {
				t.Fatalf("wanted card is %v, but got card is %v", tt.wantCard, c)
			}
			if !cmp.Equal(tt.d, tt.wantDeck, cmpopts.EquateEmpty()) {
				t.Fatalf("wanted deck is %v, but got deck is %v", tt.wantDeck.Cards, tt.d.Cards)
			}
		})
	}
}

func TestDeck_Reset_Randomness(t *testing.T) {
	t.Parallel()

	d := &Deck{}
	d.Reset()
	a := *d
	d.Reset()
	b := *d

	if cmp.Equal(a.Cards, b.Cards) {
		t.Fatalf("deck was not shuffle")
	}
}

func TestDeck_Reset_HaveAllCards(t *testing.T) {
	t.Parallel()

	checkList := map[Suit]map[CardRank]bool{
		Spade:   {},
		Club:    {},
		Diamond: {},
		Heart:   {},
	}

	d := &Deck{}
	d.Reset()

	if len(d.Cards) != 52 {
		t.Fatalf("deck not has 52 cards(%d cards)", len(d.Cards))
	}
	for _, v := range d.Cards {
		if v.Suit != Spade && v.Suit != Club && v.Suit != Diamond && v.Suit != Heart {
			t.Fatalf("invalid suit: %s", v.Suit)
		}
		if v.Rank < Deuce || Ace < v.Rank {
			t.Fatalf("invalid rank: %d", v.Rank)
		}
		_, ok := checkList[v.Suit][v.Rank]
		if ok {
			t.Fatalf("card duplicated")
		}
		checkList[v.Suit][v.Rank] = true
	}
}

func TestHand_sort(t *testing.T) {
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

			tt.hand.sort()
			if diff := cmp.Diff(tt.hand.Cards, tt.want.Cards); diff != "" {
				t.Fatalf("want and got are different(-got +want): %s", diff)
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
