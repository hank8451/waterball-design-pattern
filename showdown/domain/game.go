package domain

type game struct {}

func (g *game) Start() {}

func NewGame() *game {
	return &game{}
}