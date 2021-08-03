package lexer_test

import (
	"github.com/christiannicola/glox/internal/lexer"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func randomTestString(t *testing.T, n int, garbage bool) string {
	t.Helper()

	charset := "(){}.,-+;*"

	if garbage {
		charset = " abcdefghijklmnopqrstuvwxyz123456789\n"
	}

	var permitted = []rune(charset)

	s := make([]rune, n)

	for i := range s {
		s[i] = permitted[rand.Intn(len(permitted))]
	}

	return string(s)
}

func TestScanner_ScanTokens(t *testing.T) {
	sampleSize := 50

	scanner := lexer.NewScanner(randomTestString(t, sampleSize, false))
	tokens, err := scanner.ScanTokens()

	assert.NoError(t, err)
	// NOTE (c.nicola): We add an EOF token to the set, so len should be sampleSize + 1
	assert.Len(t, tokens, sampleSize+1)

	scanner = lexer.NewScanner(randomTestString(t, sampleSize, true))
	tokens, err = scanner.ScanTokens()

	assert.Error(t, err)
	assert.Nil(t, tokens)
}
