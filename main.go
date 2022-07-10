package main

import (
	"practice/cardGames"
)

func main() {
	players := make([]*cardGames.Player, 1)
	players[0] = cardGames.NewPlayer(cardGames.PlayerTypeHuman, "Charles", 10, cardGames.BasicBlackjackCPU{})
	cardGames.NewBlackjack(players, 1).Play()
}
