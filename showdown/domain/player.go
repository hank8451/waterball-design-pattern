package domain

import (
	"fmt"
	"math/rand"
)
const EXCHANGE_CARD_HOLD_ROUND = 3

type Player interface {
	NameMyself()
	DrawCard(deck *deck)
	addHandCard(card Card)
	DecideExchangeCard()
	exchangeCard(player *Player)
	pickExchangee() Player
}

type player struct {
	Point        int
	Name         string
	exchangeable bool
	game
	HandCard
}

type exchangeHandCard struct {
	exchangee      Player
	remainingRound int
}

func (p *player) DrawCard(deck *deck) {
	p.addHandCard(deck.DrawCard())
}

func (p *player) addHandCard(card Card) {
	p.HandCard.addCard(card)
}

func (p *player) exchangeCard(player *Player) {

}

type HumanPlayer struct {
	player
}

func (h *HumanPlayer) NameMyself() {}

func (h *HumanPlayer) DecideExchangeCard() {
	if h.exchangeable {
		// pick one player to exchange card
		exchange := exchangeHandCard{
			exchangee: h.pickExchangee(),
			remainingRound: EXCHANGE_CARD_HOLD_ROUND,
		}
		fmt.Println(exchange)
	}
}

func (h *HumanPlayer) pickExchangee() Player {
	randNumber := rand.Intn(len(h.game.players))
	return h.game.players[randNumber]
}

type AIPlayer struct {
	player
}

func (a *AIPlayer) NameMyself() {}

func (a *AIPlayer) DecideExchangeCard() {
}

func (a *AIPlayer) pickExchangee() Player {
	randNumber := rand.Intn(len(a.game.players))
	return a.game.players[randNumber]
}

func NewHumanPlayer() Player {
	return &HumanPlayer{
		player: player{},
	}
}

func NewAIPlayer() Player {
	return &AIPlayer{
		player: player{},
	}
}
