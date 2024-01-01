package domain

type Player interface {
	NameMyself()
}

type player struct {
	Point int
	Name  string
}

type HumanPlayer struct {
	player
}

func (h *HumanPlayer) NameMyself() {}

type AIPlayer struct {
	player
}

func (a *AIPlayer) NameMyself() {}

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