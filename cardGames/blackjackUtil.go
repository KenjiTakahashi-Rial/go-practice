package cardGames

import (
	"fmt"
	"strconv"
	"strings"
)

func scanMinInt(prompt string, min int) int {
	for {
		fmt.Print(prompt)
		var response string
		fmt.Scanln(&response)
		toInt, err := strconv.Atoi(response)

		if err == nil && toInt >= min {
			return toInt
		}
		fmt.Print("Invalid number. ")
	}
}

func scanPlayer(playerType PlayerType, playerNumber int) *Player {
	fmt.Printf("%s %d name?: ", playerType, playerNumber)
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	balance := scanMinInt(fmt.Sprintf("%s starting balance?: ", name), 1)
	return NewPlayer(playerType, name, balance)
}

func newBlackjackFromInput() Blackjack {
	humanCount := scanMinInt("Number of human players?: ", 1)
	humans := make([]*Player, humanCount)
	for i := range humans {
		humans[i] = scanPlayer(Human, i+1)
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
