package domain

type matchMaking struct {
	individuals []Individual
}

func NewMatchMaking() *matchMaking {
	m := &matchMaking{}
	initMatchMaking(m)
	return m
}

func initMatchMaking(m *matchMaking) {
	individuals := []Individual {
		NewIndividual(1, "MALE", 20, "Halo", []string{"Surfing", "Climbing"}, [2]float64{1.2, 2.3}, &DistanceMatchStrategy{}),
		NewIndividual(2, "FEMALE", 40, "YOLO", []string{}, [2]float64{2.2, 3.3}, &ReverseMatchStrategy{preMatch: &DistanceMatchStrategy{}}),
		NewIndividual(3, "MALE", 25, "", []string{"Shopping", "Sleeping"}, [2]float64{2.9, 9.3}, &HabitsMatchStrategy{}),
	}

	m.individuals = append(m.individuals, individuals...)
}

func (m *matchMaking) Match() {
	for _, individual := range m.individuals {
		individual.MatchStrategy.Match(individual, m.individuals)
	}
}