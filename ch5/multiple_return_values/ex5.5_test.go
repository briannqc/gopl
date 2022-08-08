package multiplereturnvalues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWordsAndImages(t *testing.T) {
	nwords, nimages, err := CountWordsAndImages("https://golang.org")
	assert.NoError(t, err)
	fmt.Println("nwords:", nwords)
	fmt.Println("nimages:", nimages)
}
