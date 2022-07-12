package run

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExisting(t *testing.T) {
	res := WhichNotInstalled([]string{"bash", "go"})
	assert.Equal(t, "", res)
}
func TestNonExisting(t *testing.T) {
	nonExisting := "non_existing_program"
	res := WhichNotInstalled([]string{"bash", nonExisting, "go"})
	assert.Equal(t, nonExisting, res)
}
