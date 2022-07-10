package cardGames

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

type BlackjackHand struct {
	hand      *Hand
	softScore int
	hardScore int
}

func NewBlackjackHand() *BlackjackHand {
	return &BlackjackHand{NewHand(), 0, 0}
}

func (h BlackjackHand) String() string { return h.hand.String() }

func (h *BlackjackHand) Reset() {
	h.hand.cards.Reset()
	h.softScore, h.hardScore = 0, 0
}

func (h *BlackjackHand) Add(c Card) {
	cardHard, cardSoft := scoreCardBlackjack(c)
	h.softScore += cardSoft
	h.hardScore += cardHard
	h.hand.Add(c)
}
