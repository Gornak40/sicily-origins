package gamecore

import "fmt"

// Last 9, 6 or 3 bits represent the player's hand.
type DeckHand int

// Turn must be less than 8.
func (d DeckHand) GetHand(turn int) []Card {
	cards := make([]Card, 0, 3)
	cards = append(cards, Card(d&7))
	if turn < 7 {
		cards = append(cards, Card(d>>3&7))
	}
	if turn < 6 {
		cards = append(cards, Card(d>>6&7))
	}

	return cards
}

// Place a card on top of the deck.
func (d *DeckHand) Append(c Card) {
	*d = *d<<3 | DeckHand(c)
}

// Remove a card with index idx (0, 1 or 2) from hand.
func (d *DeckHand) DropCard(idx int) Card {
	switch idx {
	case 0:
		return d.drop0()
	case 1:
		return d.drop1()
	case 2:
		return d.drop2()
	default:
		panic(fmt.Sprintf("unknown card index %d", idx))
	}
}

func (d *DeckHand) drop0() Card {
	c := Card(*d & 7)
	*d >>= 3

	return c
}

func (d *DeckHand) drop1() Card {
	c := Card(*d >> 3 & 7)
	c0 := *d & 7
	*d = *d>>6<<3 | c0

	return c
}

func (d *DeckHand) drop2() Card {
	c := Card(*d >> 6 & 7)
	c01 := *d & 0x3f
	*d = *d>>9<<6 | c01

	return c
}
