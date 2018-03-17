package boolean_test

import (
	"testing"
	"github.com/parth/boolean"
	"github.com/stretchr/testify/assert"
)

func TestBtoI(t *testing.T) {
	assert.Equal(t, boolean.BtoI(false), 0)
	assert.NotEqual(t, boolean.BtoI(true), 0)
}

func TestItoB(t *testing.T) {
	assert.True(t, boolean.ItoB(1))
	assert.False(t, boolean.ItoB(0))
}
