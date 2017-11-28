package jikan_test

import (
	"testing"

	"github.com/animenotifier/jikan"
	"github.com/stretchr/testify/assert"
)

func TestAnime(t *testing.T) {
	testIDs := []string{
		"1",
		"35082",
		"21557",
	}

	for _, id := range testIDs {
		anime, err := jikan.GetAnime(id)

		assert.NoError(t, err)
		assert.NotNil(t, anime)
		assert.NotEmpty(t, anime.Title)
	}
}
