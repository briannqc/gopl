package sec7_2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestNewStringReader(t *testing.T) {
	doc := `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

</body>
</html>
`
	reader := NewStringReader(doc)
	node, err := html.Parse(reader)
	assert.NoError(t, err)
	forEach(node, 0)
}

func forEach(n *html.Node, depth int) {
	fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEach(c, depth+1)
	}
}
