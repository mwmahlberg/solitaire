package solitaire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	d := []Card{}
	d = append(d, initialDeck...)
	deck := Deck{}
	// Initialize the deck with the initial order
	for i := 0; i < len(d); i++ {
		deck.order[i] = d[i]
	}

	assert.Equal(t, 0, deck.Position(), "Initial position should be 0")
	deck.Next()
	assert.Equal(t, 1, deck.Position(), "Position should be 1 after moving next")
	deck.Previous()
	assert.Equal(t, 0, deck.Position(), "Position should be 0 after moving previous")
	deck.Previous()
	assert.Equal(t, 53, deck.Position(), "Position should wrap around to 53 after moving previous")
	assert.Equal(t, deck.Value, &Card{color: jokers, card: blackJokerCard}, "Current card should be the last card in the deck")
	deck.MoveCurrent(-10)
	deck.SetPosition(43)
	assert.Equal(t, 43, deck.Position(), "Position should be set to 43")
	assert.Equal(t, deck.Value, &Card{color: jokers, card: blackJokerCard}, "Current card should be the the the joker")
	deck.MoveCurrent(15)
	assert.Equal(t, deck.Position(), 43, "Position should be 43 after moving current")
	assert.NotEqual(t, deck.Value, &Card{color: jokers, card: blackJokerCard}, "Current card should not be the joker")
	deck.SetPosition(4)
	assert.Equal(t, deck.Position(), 4, "Position should be set to 4")
	assert.Equal(t, &Card{color: jokers, card: blackJokerCard}, deck.Value, "Current card should be the joker")
}

func TestJoker(t *testing.T) {
	d := []Card{}
	d = append(d, initialDeck...)
	deck := Deck{}
	// Initialize the deck with the initial order
	for i := 0; i < len(d); i++ {
		deck.order[i] = d[i]
	}
	assert.Equal(t, 52, deck.FindRedJoker(), "Red joker should be at position 52")
	assert.Equal(t, 53, deck.FindBlackJoker(), "Black joker should be at position 53")
}

func TestSetup(t *testing.T) {
	setUp("test")
}

func TestTrueSetup(t *testing.T) {
	trueSetup("CRYPTONOMICON")
}
