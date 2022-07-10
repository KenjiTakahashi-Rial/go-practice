package cardGames

import (
	"fmt"
	"practice/scan"
	"strings"
)

const blackjackMaxBalance int = 10000000000

func scanCPULevel() cpu {
	cpuLevels := []string{
		"1. Basic",
		"2. Advanced",
	}
	prompt := fmt.Sprintf("%s\nCPU level?: ", strings.Join(cpuLevels, "\n"))
	switch scan.ScanInt(prompt, 1, len(cpuLevels)) {
	case 1:
		return BasicBlackjackCPU{}
	case 2:
		return BasicBlackjackCPU{} // TODO: Change this to advanced CPU
	default:
		panic("Invalid CPU level")
	}
}

func scanPlayer(playerType PlayerType, playerNumber int) *Player {
	fmt.Printf("%s %d name?: ", playerType, playerNumber)
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	balance := scan.ScanInt(fmt.Sprintf("%s starting balance?: ", name), 1, blackjackMaxBalance)

	var playerCPU cpu
	if playerType == PlayerTypeCPU {
		playerCPU = scanCPULevel()
	}

	return NewPlayer(playerType, name, balance, playerCPU)
}

func newBlackjackFromInput() Blackjack {
	players := make([]*Player, 0)

	for len(players) == 0 {
		humans := scan.ScanInt("Number of human players?: ", 0, 3)
		for i := 0; i < humans; i++ {
			players = append(players, scanPlayer(PlayerTypeHuman, i+1))
		}

		cpus := scan.ScanInt("Number of CPU players?: ", 0, 3)
		for i := 0; i < cpus; i++ {
			players = append(players, scanPlayer(PlayerTypeCPU, i+1))
		}

		if len(players) == 0 {
			fmt.Println("Must have at least 1 player.")
		}
	}

	minBet := scan.ScanInt("Minimum bet?: ", 1, blackjackMaxBalance)

	return NewBlackjack(players, minBet)
}

func scoreCardBlackjack(c Card) (int, int) {
	if isFace(c) {
		return 10, 10
	}

	hard := int(c.rank)
	soft := hard
	if c.rank == Ace {
		soft += 10
	}

	return hard, soft
}

func isBlackjack(h BlackjackHand) bool {
	return h.softScore == 21 && h.hand.Len() == 2
}

func finalScore(h BlackjackHand) int {
	if h.softScore <= 21 {
		return h.softScore
	}
	return h.hardScore
}

func upCard(h BlackjackHand) Card {
	for _, c := range h.hand.cards.Slice() {
		if c.faceUp {
			return c
		}
	}
	return Card{}
}

func PlayBlackjack() {
	b := newBlackjackFromInput()
	b.Play()
}
