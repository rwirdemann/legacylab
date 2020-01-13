package identation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeadingSpaces(t *testing.T) {
	assert.Equal(t, 0, LeadingSpaces("func"))
	assert.Equal(t, 2, LeadingSpaces("  func"))
	assert.Equal(t, 4, LeadingSpaces("    func"))
}
