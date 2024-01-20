package domain

type Hand struct {
	cards []Card
}

const handCardsMaxSize = 13

func (h *Hand) AddCard(card Card) {
	h.cards = append(h.cards, card)
}

func (h *Hand) PickCard(index int) Card {
	card := h.cards[index]
	h.cards = append(h.cards[:index], h.cards[index+1:]...)
	return card
}

func (h Hand) Size() int {
	return len(h.cards)
}