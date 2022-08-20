package ch7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounter_Write(t *testing.T) {
	s := "Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines. You will find bufio.ScanWords useful."
	var wordCounter WordCounter
	_, err := wordCounter.Write([]byte(s))
	assert.NoError(t, err)
	assert.Equal(t, 19, int(wordCounter))
}

func TestLineCounter_Write(t *testing.T) {
	s := `Sed ut perspiciatis unde omnis iste natus error
sit voluptatem accusantium doloremque laudantium, totam rem
aperiam, eaque ipsa quae ab illo inventore veritatis et quasi
architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam
voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed
quia consequuntur magni dolores eos qui ratione voluptatem sequi
nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor
sit amet, consectetur, adipisci velit, sed quia non numquam eius
modi tempora incidunt ut labore et dolore magnam aliquam quaerat
voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem
ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi
consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate
velit esse quam nihil molestiae consequatur, vel illum qui dolorem
eum fugiat quo voluptas nulla pariatur?`

	var lineCounter LineCounter
	_, err := lineCounter.Write([]byte(s))
	assert.NoError(t, err)
	assert.Equal(t, 14, int(lineCounter))
}
