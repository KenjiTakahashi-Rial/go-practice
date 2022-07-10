package cardGames

// TODO: Adjust for doubling and splitting when implemented

var hardTotals = [22][11]action{
	//               | hard total |
	// dealer upcard                       N/A,       A,       2,       3,       4,       5,       6,       7,       8,       9,      10,
	// --------------              -------------------------------------------------------------------------------------------------------
	/*               |          0 |*/ {invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid},
	/*               |          1 |*/ {invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid},
	/*               |          2 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          3 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          4 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          5 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          6 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          7 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          8 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          9 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         10 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         11 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         12 |*/ {invalid,     hit,     hit,     hit,   stand,   stand,   stand,     hit,     hit,     hit,     hit},
	/*               |         13 |*/ {invalid,     hit,   stand,   stand,   stand,   stand,   stand,     hit,     hit,     hit,     hit},
	/*               |         14 |*/ {invalid,     hit,   stand,   stand,   stand,   stand,   stand,     hit,     hit,     hit,     hit},
	/*               |         15 |*/ {invalid,     hit,   stand,   stand,   stand,   stand,   stand,     hit,     hit,     hit,     hit},
	/*               |         16 |*/ {invalid,     hit,   stand,   stand,   stand,   stand,   stand,     hit,     hit,     hit,     hit},
	/*               |         17 |*/ {invalid,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand},
	/*               |         18 |*/ {invalid,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand},
	/*               |         19 |*/ {invalid,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand},
	/*               |         20 |*/ {invalid,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand},
	/*               |         21 |*/ {invalid,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand},
}

var softTotals = [22][11]action{
	//               | soft total |
	// dealer upcard                       N/A,       A,       2,       3,       4,       5,       6,       7,       8,       9,      10,
	// --------------              -------------------------------------------------------------------------------------------------------
	/*               |          0 |*/ {invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid},
	/*               |          1 |*/ {invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid, invalid},
	/*               |          2 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          3 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          4 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          5 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          6 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          7 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          8 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |          9 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         10 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         11 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         12 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         13 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         14 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         15 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         16 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         17 |*/ {invalid,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit,     hit},
	/*               |         18 |*/ {invalid,     hit,   stand,   stand,   stand,   stand,   stand,   stand,   stand,     hit,     hit},
	/*               |         19 |*/ {invalid,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand},
	/*               |         20 |*/ {invalid,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand},
	/*               |         21 |*/ {invalid,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand,   stand},
}

type BasicBlackjackCPU struct{}

func (b BasicBlackjackCPU) Bet(balance, minBet int) int {
	return minBet
}

func (b BasicBlackjackCPU) Turn(hand BlackjackHand, upCard Card) action {
	dealerScore, _ := scoreCardBlackjack(upCard)

	if hand.softScore < hand.hardScore {
		return softTotals[hand.softScore][dealerScore]
	}

	return hardTotals[hand.hardScore][dealerScore]
}
