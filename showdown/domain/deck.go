package domain

import (
	"fmt"
	"showdown/domain/enum"
)

var (
	ranks = [13]enum.Rank{enum.TWO, enum.THREE, enum.FOUR, enum.FIVE, enum.SIX, enum.SEVEN, enum.EIGHT, enum.NINE, enum.TEN, enum.JACK, enum.QUEEN, enum.KING, enum.ACE}
	suits = [4]enum.Suit{enum.CLUB, enum.DIAMOND, enum.HEART, enum.SPADE}
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	initDeck(deck)
	return deck
	
}

func initDeck(deck *Deck) {
	cards := make([]Card, 0, 52)

	for _, rank := range ranks {
		for _, suit := range suits {
			cards = append(cards, NewCard(rank, suit))
		}
	}
	deck.cards = cards
}

func (d *Deck) Shuffle() {
	fmt.Println("shuffling.....")
}

func (d *Deck) DrawCard() (card Card) {
	card = d.cards[0]
	d.cards = d.cards[1:]
	return
}
