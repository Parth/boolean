package boolean_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/parth/boolean"
)

func testBtoI(t *testing.T) {
	assert.Equal(boolean.BtoI(false), 0)
	assert.NotEqual(boolean.BtoI(true), 0)
}

func testItoB(t *testing.T) {
	assert.True(1)
	assert.False(0)
}
