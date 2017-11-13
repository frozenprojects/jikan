package jikan_test

import (
	"testing"

	"github.com/animenotifier/jikan"
	"github.com/stretchr/testify/assert"
)

func TestCharacter(t *testing.T) {
	character, err := jikan.GetCharacter("1")

	assert.NoError(t, err)
	assert.NotNil(t, character)
	assert.NotEmpty(t, character.Name)
}
