package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparePasswordHash(t *testing.T) {
	t.Parallel()

	pass := "loremipsumdolor"
	got, err := HashPassword(pass)

	assert := assert.New(t)

	assert.Nilf(err, "unexpected error during HashPassword('%s')\n", pass)

	eq := ComparePasswordHash(pass, got)

	assert.True(eq, "expected to get true but got false on ComparePasswordHash()\n")
}
