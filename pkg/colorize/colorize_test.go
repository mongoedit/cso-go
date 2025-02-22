package colorize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindowsColorize(t *testing.T) {
	testCase := "this is a test of colorize"

	colorizedString := "this is a " + Red("test") + " of colorize"

	assert.Equal(t, testCase, colorizedString)
}
