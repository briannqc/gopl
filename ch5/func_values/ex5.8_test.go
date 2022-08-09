package funcvalues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindElementByID(t *testing.T) {
	doc, err := fetch("https://www.w3schools.com/html/default.asp")
	assert.NoError(t, err)

	element := FindElementByID(doc, "internalCourses")
	assert.NotNil(t, element)
}
