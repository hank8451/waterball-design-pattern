package domain

import (
	"sort"
)
type (
	DistanceMatchStrategy struct {}
	HabitsMatchStrategy struct {}
	ReverseMatchStrategy struct {
		preMatch MatchStrategy
	}
)

type MatchStrategy interface {
	Match(Individual, []Individual) []Individual
}

func (s *DistanceMatchStrategy) Match(self Individual, individuals []Individual) []Individual {
	sort.Slice(individuals, func(i, j int) bool {
		return self.DistanceBetween(individuals[i]) < self.DistanceBetween(individuals[j])
	})
	return individuals
}

func (s *HabitsMatchStrategy) Match(self Individual, individuals []Individual) []Individual {
	sort.Slice(individuals, func(i, j int) bool {
		return len(self.HabitsIntersection(individuals[i])) > len(self.HabitsIntersection(individuals[j]))
	})
	return individuals
}

func (s *ReverseMatchStrategy) Match(self Individual, individuals []Individual) []Individual {
	individuals = s.preMatch.Match(self, individuals)
	reverse := []Individual{}
	for i := len(individuals) - 1; i >= 0; i-- {
		reverse = append(reverse, individuals[i])
	}
	return reverse
}