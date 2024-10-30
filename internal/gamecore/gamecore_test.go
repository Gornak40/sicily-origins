package gamecore_test

import (
	"testing"

	"github.com/Gornak40/sicily-origins/internal/gamecore"
	"github.com/stretchr/testify/assert"
)

func TestAppendGetHand(t *testing.T) {
	t.Parallel()

	var d gamecore.DeckHand
	d.Append(gamecore.CardKiller)
	d.Append(gamecore.CardJoker)
	d.Append(gamecore.CardCapo)
	d.Append(gamecore.CardAce)

	assert.Equal(t, []gamecore.Card{gamecore.CardAce, gamecore.CardCapo, gamecore.CardJoker}, d.GetHand(0))
	assert.Equal(t, []gamecore.Card{gamecore.CardAce, gamecore.CardCapo, gamecore.CardJoker}, d.GetHand(5))
	assert.Equal(t, []gamecore.Card{gamecore.CardAce, gamecore.CardCapo}, d.GetHand(6))
	assert.Equal(t, []gamecore.Card{gamecore.CardAce}, d.GetHand(7))
}

func TestAppendDrop(t *testing.T) {
	t.Parallel()

	// Russian deck (two killers).
	var d gamecore.DeckHand
	d.Append(gamecore.CardAce)
	d.Append(gamecore.CardDon)
	d.Append(gamecore.CardJoker)
	d.Append(gamecore.CardSoldier)
	d.Append(gamecore.CardKiller)
	d.Append(gamecore.CardGang)
	d.Append(gamecore.CardOfficer)
	d.Append(gamecore.CardKiller)

	assert.Equal(t, gamecore.CardKiller, d.DropCard(0))
	assert.Equal(t, gamecore.CardGang, d.DropCard(1))
	assert.Equal(t, gamecore.CardOfficer, d.DropCard(0))
	assert.Equal(t, gamecore.CardJoker, d.DropCard(2))
	assert.Equal(t, gamecore.CardSoldier, d.DropCard(1))
	assert.Equal(t, gamecore.CardKiller, d.DropCard(0))
	assert.Equal(t, gamecore.CardAce, d.DropCard(1))
	assert.Equal(t, gamecore.CardDon, d.DropCard(0))

	assert.Panics(t, func() { d.DropCard(3) })
	assert.Panics(t, func() { d.DropCard(-1) })
}

func TestGetName(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { gamecore.Card(8).GetName() })
	assert.Panics(t, func() { gamecore.Card(-1).GetName() })
}
