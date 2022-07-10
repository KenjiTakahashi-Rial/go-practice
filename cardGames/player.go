package cardGames

import "fmt"

type PlayerType int

const (
	PlayerTypeHuman PlayerType = iota
	PlayerTypeCPU
	PlayerTypeDealer
)

func (p PlayerType) String() string {
	switch p {
	case PlayerTypeHuman:
		return "Human"
	case PlayerTypeCPU:
		return "CPU"
	case PlayerTypeDealer:
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
	playerCPU  cpu
}

func NewPlayer(playerType PlayerType, name string, balance int, playerCPU cpu) *Player {
	return &Player{playerType, name, NewBlackjackHand(), balance, playerCPU}
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
