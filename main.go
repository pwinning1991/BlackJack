package main

import (
	"strings"

	"github.com/pwinning1991/deck"
)

type Hand []deck.Card

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDDEN**"
}

func (h Hand) MinScore() int {
	runningTotal := 0
	for _, card := range h {
		runningTotal += Min(int(card.Rank), 10)
	}
	return runningTotal
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}

	for _, card := range h {
		if card.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func main() {
	//cards := deck.New(deck.Deck(3), deck.Shuffle)
	var gs GameState
	gs.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	//var card deck.Card
	//var player, dealer Hand
	//for i := 0; i < 2; i++ {
	//for _, hand := range []*Hand{&player, &dealer} {
	//card, cards = draw(cards)
	//*hand = append(*hand, card)
	//}
	//}
	//var input string
	//for input != "s" {
	//fmt.Println("Player:", player)
	//fmt.Println("Player Score:", player.Score())
	//fmt.Println("Dealer:", dealer.DealerString())
	//fmt.Println("What will you do? (h)it, (s)tand")
	//fmt.Scanf("%s\n", &input)
	//switch input {
	//case "h":
	//card, cards = draw(cards)
	//player = append(player, card)
	//}
	//}
	//for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
	//card, cards = draw(cards)
	//dealer = append(dealer, card)
	//}
	//pScore, dScore := player.Score(), dealer.Score()
	//fmt.Println("==FINAL HANDS==")
	//fmt.Println("Player:", player)
	//fmt.Println("Dealer:", dealer)
	//fmt.Println("Player Score:", pScore)
	//fmt.Println("Dealer Score:", dScore)
	//switch {
	//case pScore > 21:
	//fmt.Println("You Busted")
	//case dScore > 21:
	//fmt.Println("Dealer Busted")
	//case pScore > dScore:
	//fmt.Println("You Win")
	//case pScore < dScore:
	//fmt.Println("You Lose")
	//case pScore == dScore:
	//fmt.Println("Draw")

	//}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("it isn't currently any player's turn")
	}

}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}
