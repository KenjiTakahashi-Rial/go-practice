package cardGames

import (
	"strings"

	"practice/collections"
	"practice/functional"
)

type Hand struct {
	cards *collections.LinkedList[Card]
}

func NewHand() *Hand {
	return &Hand{collections.NewLinkedList[Card]()}
}

func (h Hand) Len() int                   { return h.cards.Len() }
func (h Hand) Add(c Card)                 { h.cards.PushBack(c) }
func (h Hand) Remove(c Card) (Card, bool) { return collections.RemoveFirst(h.cards, c) }

func (h Hand) String() string {
	symbol := func(c Card) string {
		return c.Symbol()
	}
	cardSymbols := functional.Map(h.cards.Slice(), symbol)
	return strings.Join(cardSymbols, "")
}
