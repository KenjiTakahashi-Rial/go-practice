package cardGames

import (
	"fmt"
	"strconv"
	"strings"

	"practice/collections"
)

type action int

const (
	hit action = iota + 1
	stand
	double
	split
)

type Blackjack struct {
	dealer  *Player
	players *collections.LinkedList[*Player]
	deck    *Deck
	minBet  int
	bets    map[*Player]int
}

func NewBlackjack(humans, cpus []*Player, minBet int) Blackjack {
	dealer := NewPlayer(PlayerTypeDealer, PlayerTypeDealer.String(), 0, nil)
	newPlayers := collections.NewLinkedList[*Player]()

	for _, p := range humans {
		newPlayers.PushBack(p)
	}
	for _, p := range cpus {
		newPlayers.PushBack(p)
	}

	deck := NewDeck()
	deck.Shuffle()
	bets := make(map[*Player]int, newPlayers.Len())
	return Blackjack{dealer, newPlayers, deck, minBet, bets}
}

func (b Blackjack) acceptBets() {
	for _, p := range b.players.Slice() {
		for {
			fmt.Printf("Bet for %s (balance %d) or 0 to quit: ", p.name, p.balance) // TODO: Format numbers over 1000 with commas and add dollar signs
			var response string
			fmt.Scanln(&response)
			bet, err := strconv.Atoi(response)

			if err != nil {
				fmt.Print("Invalid bet. ")
			} else if bet == 0 {
				collections.RemoveFirst(b.players, p)
				break
			} else if bet < b.minBet {
				fmt.Printf("Minimum bet is %d. ", b.minBet)
			} else if !p.SubtractBalance(bet) {
				fmt.Print("Bet exceeds balance. ")
			} else {
				b.bets[p] = bet
				break
			}
		}
	}
}

func (b Blackjack) deal(player Player, faceUp bool) {
	if b.deck.Len() == 0 {
		b.deck.Reset()
		b.deck.Shuffle()
	}

	top := b.deck.Top()
	top.faceUp = faceUp

	player.hand.Add(top)
}

func (b Blackjack) dealFirstHand() {
	players := b.players.Slice()

	for _, p := range players {
		b.deal(*p, p.playerType != PlayerTypeDealer)
	}
	b.deal(*b.dealer, true)

	for _, p := range players {
		b.deal(*p, true)
	}
	b.deal(*b.dealer, false)
}

func (b Blackjack) handleDealerBlackjack() {
	fmt.Println("Dealer Blackjack.")

	for _, p := range b.players.Slice() {
		if isBlackjack(*p.hand) {
			fmt.Printf("%s Blackjack. Push.\n", p.name)
		} else {
			fmt.Printf("%s %s. Lose.\n", p.name, p.hand)
		}
	}
}

func (b Blackjack) acceptAction(p Player) action {
	actionStrs := []string{
		"1. Hit",
		"2. Stand",
	}
	prompt := fmt.Sprintf("%s\nAction?: ", strings.Join(actionStrs, "\n"))
	fmt.Print(prompt)

	for {
		var response string
		fmt.Scanln(&response)
		a, err := strconv.Atoi(response)

		if err != nil {
			fmt.Printf("Invalid action. %s", prompt)
		} else if a < 1 || a > len(actionStrs) {
			fmt.Printf("Action must be between 1 and %d. %s", len(actionStrs), prompt)
		} else {
			return action(a)
		}
	}
}

func (b Blackjack) handleAction(p Player, a action) {
	switch a {
	case hit:
		b.deal(p, true)
	case stand:
		break
	case double:
		// TODO
	case split:
		// TODO
	}
}

func (b Blackjack) printHands(players ...Player) {
	for _, p := range players {
		fmt.Printf("%s's hand: %s\n", p.name, p.hand)
	}
}

func (b Blackjack) playerTurn(p Player) {
	b.printHands(p, *b.dealer)
	for p.hand.hardScore <= 21 {
		a := b.acceptAction(p)
		b.handleAction(p, a)
		if a == stand {
			break
		}
		b.printHands(p, *b.dealer)
	}
}

func (b Blackjack) playerTurns() []*Player {
	remainingPlayers := make([]*Player, 0)
	for _, p := range b.players.Slice() {
		if isBlackjack(*p.hand) {
			b.printHands(*p, *b.dealer)
			b.win(p)
		}

		b.playerTurn(*p) // TODO: CPU turns
		if p.hand.hardScore > 21 {
			b.bust(p)
		} else {
			remainingPlayers = append(remainingPlayers, p)
		}

		fmt.Println()
	}
	return remainingPlayers
}

func (b Blackjack) dealerTurn() {
	for _, c := range b.dealer.hand.hand.cards.PointerSlice() {
		c.faceUp = true
	}
	fmt.Printf("Dealer's turn.\n%s\n", b.dealer.hand)

	for b.dealer.hand.softScore < 17 || b.dealer.hand.softScore > 21 && b.dealer.hand.hardScore < 17 {
		b.deal(*b.dealer, true)
		fmt.Println(b.dealer.hand)
	}
}

func (b Blackjack) win(p *Player) {
	bet := b.bets[p]
	blackjackStr := ""

	if isBlackjack(*p.hand) {
		blackjackStr = "Blackjack! "
	}

	fmt.Printf("%s%s wins $%d!\n", blackjackStr, p.name, bet)
	p.AddBalance(2 * bet)
}

func (b Blackjack) lose(p *Player) {
	fmt.Printf("%s loses.\n", p.name)
}

func (b Blackjack) bust(p *Player) {
	fmt.Printf("%s bust.\n", p.name)
}

func (b Blackjack) push(p *Player) {
	p.AddBalance(b.bets[p])
	fmt.Printf("%s push.\n", p.name)
}

func (b Blackjack) judgeRemaining(remainingPlayers []*Player) {
	if b.dealer.hand.hardScore > 21 {
		for _, p := range remainingPlayers {
			fmt.Println("Dealer bust.")
			b.win(p)
		}
	} else {
		dealerScore := finalScore(*b.dealer.hand)
		for _, p := range remainingPlayers {
			b.printHands(*p)
			playerScore := finalScore(*p.hand)

			if playerScore < dealerScore {
				b.lose(p)
			} else if playerScore > dealerScore {
				b.win(p)
			} else {
				b.push(p)
			}

			fmt.Println()
		}
	}
}

func (b Blackjack) checkBalances() {
	for _, p := range b.players.Slice() {
		if p.balance >= b.minBet {
			return
		}

		warning := fmt.Sprintf("Min bet (%d) exceeds %s balance (%d). ", b.minBet, p.name, p.balance)
		fmt.Print(warning)
		for {
			fmt.Printf("Enter buy-in amount or 0 to quit: ")
			var response string
			fmt.Scan(&response)
			amount, err := strconv.Atoi(response)

			if err != nil || amount < 0 {
				fmt.Print("Invalid amount. ")
			} else if amount == 0 {
				collections.RemoveFirst(b.players, p)
				break
			} else {
				p.AddBalance(amount)
				if p.balance < b.minBet {
					fmt.Print(warning)
					continue
				} else {
					break
				}
			}
		}
	}
}

func (b Blackjack) round() {
	fmt.Println("Round start.")
	defer fmt.Println("Round end.")
	b.checkBalances()
	b.acceptBets()
	b.dealFirstHand()

	if isBlackjack(*b.dealer.hand) {
		b.handleDealerBlackjack()
		return
	}

	remainingPlayers := b.playerTurns()

	if len(remainingPlayers) > 0 {
		b.dealerTurn()
		b.judgeRemaining(remainingPlayers)
	}
}
