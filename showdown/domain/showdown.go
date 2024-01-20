package domain

import "fmt"

type Showdown struct {
	players []Player
	deck    *Deck
	round   int
}

const maxRound = 13

func NewShowdown() Showdown {
	s := &Showdown{}
	initShowdown(s)
	return *s
}

func initShowdown(s *Showdown) {
	deck := NewDeck()
	s.deck = deck
}

func (s *Showdown) Start() {
	for _, player := range s.players {
		player.NameMyself()
	}

	s.deck.Shuffle()

	for _, player := range s.players {
		for handCards := player.GetHands(); handCards.Size() <= handCardsMaxSize; {
			card := s.deck.DrawCard()
			player.AddHand(card)
		}
	}

	for s.round < maxRound {
		for _, player := range s.players {
			player.DecideChangeHandCards()
		}

		showCardMap := make(map[Card]Player, s.PlayersCount())
		for _, player := range s.players {
			card := player.PlayCard()
			showCardMap[card] = player
		}

		var roundWinner Player
		var biggerCard Card
		for card, player := range showCardMap {
			if roundWinner == nil {
				roundWinner = player
				biggerCard = card
				continue
			}

			biggerCard = biggerCard.Bigger(card)
		}
		showCardMap[biggerCard].AddOnePoint()

		s.round++
	}

	var winner Player
	for _, player := range s.players {
		if winner == nil {
			winner = player
			continue
		}

		if player.GetPoint() > winner.GetPoint() {
			winner = player
		}
	}

	fmt.Printf("Winner is %s", winner.GetName())
}

func (s Showdown) PlayersCount() int {
	return len(s.players)
}

func (s *Showdown) AddPlayer(player Player) {
	player.SetGame(s)
	s.players = append(s.players, player)
}

func (s *Showdown) GetPlayers() []Player {
	return s.players
}