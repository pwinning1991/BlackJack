package main

import (
	"fmt"
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

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}

	ret.State = StatePlayerTurn
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret

}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		Stand(ret)
	}
	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", ret.Player)
	fmt.Println("Dealer:", ret.Dealer)
	fmt.Println("Player Score:", pScore)
	fmt.Println("Dealer Score:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You Busted")
	case dScore > 21:
		fmt.Println("Dealer Busted")
	case pScore > dScore:
		fmt.Println("You Win")
	case pScore < dScore:
		fmt.Println("You Lose")
	case pScore == dScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	var gs GameState
	gs = Shuffle(gs)
	for i := 0; i < 10; i++ {
		gs = Deal(gs)
		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("Player:", gs.Player)
			fmt.Println("Player Score:", gs.Player.Score())
			fmt.Println("Dealer:", gs.Dealer.DealerString())
			fmt.Println("What will you do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid option: chose either \"h\" or \"s\"")
			}

		}

		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}

		gs = EndHand(gs)
	}
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
