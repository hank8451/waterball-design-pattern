package domain

const HAND_CARD_MAX_NUM = 13
const MAX_ROUND = 13

type game struct {
	players []Player
	deck    *deck
	round   int
}

func (g *game) Start() {
	for _, player := range g.players {
		player.NameMyself()
	}

	g.deck.Shuffle()

	for i := 0; i < HAND_CARD_MAX_NUM; i++ {
		for _, player := range g.players {
			player.DrawCard(g.deck)
		}
	}

	for g.round < MAX_ROUND {
		for _, player := range g.players {
			player.DecideExchangeCard()
		}
		g.round++
	}

}

func (g *game) AddPlayer(player Player) {
	g.players = append(g.players, player)
}

func NewGame() *game {
	return &game{
		players: []Player{},
		deck:    NewDeck(),
	}
}
