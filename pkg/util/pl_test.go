package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPLMap1(t *testing.T) {
	lang, err := ExtensionToLanguge("c")
	assert.NoError(t, err)
	assert.Equal(t, "c", lang)
}

func TestPLMap2(t *testing.T) {
	lang, err := ExtensionToLanguge("ex")
	assert.NoError(t, err)
	assert.Equal(t, "elixir", lang)
}

func TestAutoLanguage(t *testing.T) {
	lang, err := AutoDetectLanguage(".")
	assert.NoError(t, err)
	assert.Equal(t, "go", lang)
}
