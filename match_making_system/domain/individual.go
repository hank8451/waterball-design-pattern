package domain

import "math"

type (
	Individual struct {
		ID            uint32
		Gender        string
		Age           uint8
		Intro         string
		Habits        []string
		Coord         coord
		MatchStrategy MatchStrategy
	}

	coord struct {
		X float64
		Y float64
	}
)

func NewIndividual(
	id uint32,
	gender string,
	age uint8,
	intro string,
	habits []string,
	position [2]float64,
	matchStrategy MatchStrategy,
) Individual {
	return Individual{
		ID:     id,
		Gender: gender,
		Age:    age,
		Intro:  intro,
		Habits: habits,
		Coord:  coord{X: position[0], Y: position[1]},
		MatchStrategy: matchStrategy,
	}
}

func (i *Individual) DistanceBetween(other Individual) float64 {
	return math.Sqrt(math.Pow(i.Coord.X-other.Coord.X, 2) + math.Pow(i.Coord.Y-other.Coord.Y, 2))
}

func (i *Individual) HabitsIntersection(other Individual) (inter []string) {
	habitMap := make(map[string]bool)
	for _, habit := range i.Habits {
		habitMap[habit] = true
	}

	for _, habit := range other.Habits {
		if habitMap[habit] {
			inter = append(inter, habit)
		}
	}
	return
}
