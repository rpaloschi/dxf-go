package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatEquals(t *testing.T) {
	tests := []struct {
		v1     float64
		v2     float64
		equals bool
	}{
		{1.1, 1.1, true},
		{1.001, 1.0009, true},
		{1.1, 2.0, false},
	}

	for _, test := range tests {
		assert.Equal(t, FloatEquals(test.v1, test.v2), test.equals,
			"Test case: %+v", test)
	}
}
