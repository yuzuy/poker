package poker

import "math/rand"

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
	ranks := []CardRank{Ace, Deuce, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
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
