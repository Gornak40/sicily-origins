package gamecore

import (
	"fmt"
)

// Card power is equal to the card value.
type Card int

const (
	CardJoker Card = iota
	CardAce
	CardGang
	CardOfficer
	CardSoldier
	CardKiller
	CardCapo
	CardDon
)

func (c Card) GetName() string {
	switch c {
	case CardJoker:
		return "Joker"
	case CardAce:
		return "Ace"
	case CardGang:
		return "Gang"
	case CardOfficer:
		return "Officer"
	case CardSoldier:
		return "Soldier"
	case CardKiller:
		return "Killer"
	case CardCapo:
		return "Capo"
	case CardDon:
		return "Don"
	default:
		panic(fmt.Sprintf("unknown card %d", c))
	}
}
