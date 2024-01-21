package main

import (
	"fmt"
	"github.com/hank8451/match_making_system/domain"
)

func main() {
	fmt.Println("Match Making System")
	matchMakingSystem := domain.NewMatchMaking()
	matchMakingSystem.Match()
}