package domain

import "fmt"

type ExchangeHands struct {
	remainingRound uint8
	exchanger      Player
	exchangee      Player
}

const rollbackRound uint8 = 3

func NewExchangeHandCards(exchanger, exchangee Player) *ExchangeHands {
	return &ExchangeHands{
		remainingRound: rollbackRound,
		exchanger: exchanger,
		exchangee: exchangee,
	}
}

func (e *ExchangeHands) Exchange() {
	fmt.Printf("%s change his handcard to %s", e.exchanger.GetName(), e.exchangee.GetName())
	p1Hand, p2Hand := e.exchanger.GetHands(), e.exchangee.GetHands()
	e.exchanger.SetHands(p2Hand)
	e.exchangee.SetHands(p1Hand)
}

func (e *ExchangeHands) rollback() {}

func (e *ExchangeHands) Countdown() {
	e.remainingRound--
	if e.remainingRound == 0 {
		e.rollback()
	}
}
