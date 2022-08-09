package funcvalues

import (
	"crypto/sha1"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	f := func(s string) string {
		hasher := sha1.New()
		hasher.Write([]byte(s))
		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		return sha[:8]
	}

	s := `The standard Lorem Ipsum passage, used since the 1500s

	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris $foo nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt $foo mollit anim id est $foo laborum.`

	want := `The standard Lorem Ipsum passage, used since the 1500s

	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris C-7Hteo_ nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt C-7Hteo_ mollit anim id est C-7Hteo_ laborum.`

	got := expand(s, f)

	assert.Equal(t, want, got)
}
