package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoAllArgsHasText(t *testing.T) {
	b := DoAllArgsHasText(" ")
	assert.True(t, !b)

	b = DoAllArgsHasText("aa")
	assert.True(t, b)
}
