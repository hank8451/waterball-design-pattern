package domain

import (
	"showdown/domain/enum"
)

type Card struct {
	rank enum.Rank
	suit enum.Suit
}

func NewCard(rank enum.Rank, suit enum.Suit) Card {
	return Card {
		rank, suit,
	}
}

func (c Card) Bigger(card Card) Card {
	if c.rank == card.rank {
		if c.suit > card.suit {
			return c
		} else {
			return card
		}
	} else if c.rank > card.rank {
		return c
	} else {
		return card
	}
}