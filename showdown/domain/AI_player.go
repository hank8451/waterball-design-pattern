package domain

import (
	"fmt"
	"math/rand"
)

type AIPlayer struct {
	player
}

func NewAIPlayer() Player {
	return &AIPlayer{
		player: player{
			canExchangeCard: true,
		},
	}
}

func (a *AIPlayer) NameMyself() {
	fmt.Println("Name randomly")
}

func (a *AIPlayer) PlayCard() Card {
	fmt.Println("pick a card randomly")
	randomIndex := rand.Intn(a.hand.Size())
	card := a.hand.PickCard(randomIndex)
	return card
}

func (a *AIPlayer) DecideChangeHandCards() {}