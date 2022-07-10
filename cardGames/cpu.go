package cardGames

type cpu interface {
	Bet(balance, minBet int) int
	Turn(hand BlackjackHand, upCard Card) action
}
