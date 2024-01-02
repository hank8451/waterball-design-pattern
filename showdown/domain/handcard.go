package domain

type HandCard struct {
	cards []Card
}

func (h *HandCard) addCard(card Card) {
	h.cards = append(h.cards, card)
}

func (h HandCard) Size() int {
	return len(h.cards)
}