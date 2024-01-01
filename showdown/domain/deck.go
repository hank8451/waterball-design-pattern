package domain

var (
	Ranks = []int{TWO, THREE, FOUR, FIVE, SIX, EIGHT, NINE, TEN, JACK, QUEEN, KING, ACE}
	Suits = []int{CLUB, DIAMOND, HEART, SPADE}
)

type deck struct {
	cards []Card
}

func NewDeck() *deck {
	cards := make([]Card, 0, 52)
	for _, s := range Suits {
		for _, r := range Ranks {
			cards = append(cards, Card{Rank: r, Suit: s})
		}
	}
	return &deck{cards: cards}
}

func (d *deck) Shuffle() {}

func (d *deck) DrawCard() (card Card) {
	card = d.cards[0]
	d.cards = d.cards[1:]
	return
}
