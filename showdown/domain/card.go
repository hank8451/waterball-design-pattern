package domain

type Card struct {
	Rank int
	Suit int
}

const (
	TWO = iota + 1
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

const (
	CLUB = iota + 1
	DIAMOND
	HEART
	SPADE
)
