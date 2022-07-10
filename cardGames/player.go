package cardGames

import "fmt"

type PlayerType int

const (
	Human PlayerType = iota
	CPU
	Dealer
)

func (p PlayerType) String() string {
	switch p {
	case Human:
		return "Human"
	case CPU:
		return "CPU"
	case Dealer:
		return "Dealer"
	default:
		return fmt.Sprintf("%d", int(p))
	}
}

type Player struct {
	playerType PlayerType
	name       string
	hand       *BlackjackHand
	balance    int
}

func NewPlayer(playerType PlayerType, name string, balance int) *Player {
	return &Player{playerType, name, NewBlackjackHand(), balance}
}

func (b *Player) SubtractBalance(amount int) bool {
	if amount > b.balance {
		return false
	}
	b.balance -= amount
	return true
}

func (b *Player) AddBalance(amount int) {
	b.balance += amount
}