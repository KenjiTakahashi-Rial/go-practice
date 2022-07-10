package cardGames

import (
	"practice/algorithms"
	"practice/functional"
	"strings"
)

type Deck struct {
	cards []*Card
}

func NewDeck() *Deck {
	d := &Deck{}
	d.Reset()
	return d
}

func (d Deck) Len() int { return len(d.cards) }

func (d Deck) String() string {
	toString := func(c *Card) string {
		return c.String()
	}
	cardStrings := functional.Map(d.cards, toString)
	return strings.Join(cardStrings, "\n")
}

func (d *Deck) Reset() {
	newCards := make([]*Card, 52)
	for i, suit := 0, Clubs; suit <= Spades; suit++ {
		for rank := Ace; rank <= King; i, rank = i+1, rank+1 {
			newCards[i] = NewCard(rank, suit, false)
		}
	}
	d.cards = newCards
}

func (d *Deck) Top() Card {
	top := d.cards[0]
	d.cards = d.cards[1:]
	return *top
}

func (d *Deck) Shuffle() {
	algorithms.Shuffle(d.cards)
}
