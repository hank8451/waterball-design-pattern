package domain

import "fmt"

type Player interface {
	GetName() string
	NameMyself()
	AddHand(Card)
	GetHands() *Hand
	SetHands(*Hand)
	DecideChangeHandCards()
	PlayCard() Card
	GetPoint() int
	AddOnePoint()
	GetGame() *Showdown
	SetGame(*Showdown)
}

type player struct {
	name            string
	point           int
	canExchangeCard bool
	game            *Showdown
	hand            *Hand
	exchangeHand    *ExchangeHands
}

func (p player) GetName() string {
	return p.name
}

func (p *player) AddHand(card Card) {
	p.hand.AddCard(card)
}

func (p *player) GetHands() *Hand {
	return p.hand
}

func (p *player) SetHands(hand *Hand) {
	p.hand = hand
}

func (p *player) GetPoint() int {
	return p.point
}

func (p *player) AddOnePoint() {
	p.point++
}

func (p *player) GetGame() *Showdown {
	return p.game
}

func (p *player) SetGame(game *Showdown) {
	p.game = game
}

func (p *player) String() string {
	return fmt.Sprintf("Name: %s", p.GetName())
}
