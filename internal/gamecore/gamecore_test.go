package gamecore_test

import (
	"testing"

	"github.com/Gornak40/sicily-origins/internal/gamecore"
	"github.com/stretchr/testify/assert"
)

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
}

func TestDropPanic(t *testing.T) {
	t.Parallel()

	var d gamecore.DeckHand
	assert.Panics(t, func() { d.DropCard(3) })
	assert.Panics(t, func() { d.DropCard(-1) })
}
