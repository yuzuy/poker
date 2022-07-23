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
