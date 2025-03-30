package solitaire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoker(t *testing.T) {
	d := []Card{}
	d = append(d, initialDeck...)
	deck := Deck{}
	// Initialize the deck with the initial order
	for i := 0; i < len(d); i++ {
		deck.order[i] = d[i]
	}
	assert.Equal(t, 52, deck.FindJokerA(), "Joker A should be at position 52")
	assert.Equal(t, 53, deck.FindJokerB(), "Joker B should be at position 53")
}

func TestMove(t *testing.T) {
	d := Deck{}
	// Initialize the deck with the initial order
	copy(d.order[:], initialDeck)
	jokerA := d.FindJokerA()
	assert.Equal(t, 52, jokerA, "Joker A should be at position 52")
	jokerB := d.FindJokerB()
	assert.Equal(t, 53, jokerB, "Joker B should be at position 53")

	d.Move(jokerA, 1)
	jokerAafterRound1 := d.FindJokerA()
	assert.Equal(t, 53, jokerAafterRound1, "Joker A should be at position 53 after round 1")
	jokerBafterRound1 := d.FindJokerB()
	assert.Equal(t, 52, jokerBafterRound1, "Joker B should be at position 52 after round 1")

	d.Move(jokerBafterRound1, 2)
	jokerAafterRound2 := d.FindJokerA()
	assert.Equal(t, 53, jokerAafterRound2, "Joker A should be at position 52 after round 2")
	jokerBafterRound2 := d.FindJokerB()
	assert.Equal(t, 1, jokerBafterRound2, "Joker B should be at position 1 after round 2")
	assert.Equal(t, 1, d.order[0].Value())
	assert.Equal(t, clubs, d.order[0].color)
	assert.Equal(t, 1, int(d.order[0].card))
}
