package cardGames

import (
	"fmt"
	"practice/scan"
	"strings"
)

const blackjackMaxBalance int = 10000000000

func scanPlayer(playerType PlayerType, playerNumber int) *Player {
	fmt.Printf("%s %d name?: ", playerType, playerNumber)
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	balance := scan.ScanInt(fmt.Sprintf("%s starting balance?: ", name), 1, blackjackMaxBalance)
	return NewPlayer(playerType, name, balance)
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

	cpuCount := scanMinInt("Number of CPU players?: ", 0)
	cpus := make([]*Player, cpuCount)
	for i := range cpus {
		cpus[i] = scanPlayer(CPU, i+1)
	}

	minBet := scanMinInt("Minimum bet?: ", 1)

	return NewBlackjack(humans, cpus, minBet)
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

func PlayBlackjack() {
	b := newBlackjackFromInput()
	for b.players.Len() > 0 {
		b.round()
		for _, p := range b.players.Slice() {
			p.hand.Reset()
		}
		b.dealer.hand.Reset()
	}
}
