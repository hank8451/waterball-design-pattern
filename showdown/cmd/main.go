package main

import (
	"showdown/domain"
)

func main() {
	game := domain.NewShowdown()

	game.AddPlayer(domain.NewHumanPlayer())
	game.AddPlayer(domain.NewAIPlayer())
	game.AddPlayer(domain.NewAIPlayer())
	game.AddPlayer(domain.NewAIPlayer())

	game.Start()
}