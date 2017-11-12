package jikan_test

import (
	"testing"

	"github.com/animenotifier/jikan"
	"github.com/stretchr/testify/assert"
)

func TestAnime(t *testing.T) {
	anime, err := jikan.GetAnime("1")
	assert.NoError(t, err)
	assert.NotNil(t, anime)
}
