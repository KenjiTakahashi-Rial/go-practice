package cardGames

import "fmt"

type Rank int

const (
	Ace Rank = iota + 1
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

func (r Rank) String() string {
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

func (r Rank) Symbol() string {
	switch r {
	case Ace:
		return "A"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}

func (s Suit) Symbol() string {
	switch s {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	default:
		return "?"
	}
}

func isFace(c Card) bool {
	switch c.rank {
	case Jack, Queen, King:
		return true
	default:
		return false
	}
}

type Card struct {
	rank   Rank
	suit   Suit
	faceUp bool
}

func NewCard(rank Rank, suit Suit, faceUp bool) *Card {
	return &Card{rank, suit, faceUp}
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.rank, c.suit)
}

func (c Card) Symbol() string {
	if c.faceUp {
		return fmt.Sprintf("[%s%s]", c.rank.Symbol(), c.suit.Symbol())
	}
	return "[?]"
}
