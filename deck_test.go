package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	assert.Equal(t, len(d), 16)
	assert.Equal(t, d[0], "Ace of Spades")
	assert.Equal(t, d[len(d)-1], "Four of Clubs")
}
