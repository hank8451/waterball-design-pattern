package domain

import (
	"fmt"
	"math/rand"
	"strconv"
)

type HumanPlayer struct {
	player
}

func NewHumanPlayer() Player {
	return &HumanPlayer{
		player: player{
			canExchangeCard: true,
		},
	}
}

func (h *HumanPlayer) NameMyself() {
	fmt.Println("please type your name....")
}

func (h *HumanPlayer) PlayCard() Card {
	fmt.Println("pick one card from hand...")
	//TODO
	randomIndex := rand.Intn(h.hand.Size())
	card := h.hand.PickCard(randomIndex)
	return card
}

func (h *HumanPlayer) DecideChangeHandCards() {
	if exchangeHand := h.exchangeHand; exchangeHand != nil {
		exchangeHand.Countdown()
		return
	}

	if !h.canExchangeCard {
		return
	}

	fmt.Print("Choose to exchange hand cards Y/N")
	var whetherExchangeHands string
	fmt.Scan(&whetherExchangeHands)
	if whetherExchangeHands != "N" && whetherExchangeHands != "Y" {
		fmt.Println("Please Enter Y or N")
		// h.DecideChangeHandCards()
		return
	} else if whetherExchangeHands == "N" {
		return
	}

	exchangeOptions := h.GetGame().GetPlayers()
	fmt.Println("Choose whose players' hands you want to exchange:")
	selfIndex := 0
	for i, player := range exchangeOptions {
		if player == h {
			selfIndex = i
			continue
		}
		fmt.Printf("%d: %s", i, player.GetName())
	}
	var exchangeeIndexString string
	fmt.Scan(&exchangeeIndexString)

	exchangeeIndex, err := strconv.ParseInt(exchangeeIndexString, 10, 0)
	if err != nil || (exchangeeIndex > 3 || exchangeeIndex < 0 || exchangeeIndex == int64(selfIndex)) {
		fmt.Println("Please Enter Correct Player Number")
		return
	}
	exchangeHand := NewExchangeHandCards(h, exchangeOptions[exchangeeIndex])
	exchangeHand.Exchange()
	h.canExchangeCard = false

}
