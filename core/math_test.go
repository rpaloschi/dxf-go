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
		{1.0000000001, 1.00000000009, true},
		{1.1, 2.0, false},
	}

	for _, test := range tests {
		assert.Equal(t, FloatEquals(test.v1, test.v2), test.equals,
			"Test case: %+v", test)
	}
}

func TestFloatSliceEquals(t *testing.T) {
	tests := []struct {
		v1     []float64
		v2     []float64
		equals bool
	}{
		{[]float64{}, []float64{}, true},
		{[]float64{1.1}, []float64{}, false},
		{[]float64{2.0}, []float64{2.0}, true},
		{[]float64{2.0}, []float64{32.2}, false},
		{[]float64{1.9, 2.0}, []float64{1.9, 32.2}, false},
		{[]float64{2.0, 109.1, 33.09}, []float64{2.0, 109.1, 33.09}, true},
	}

	for _, test := range tests {
		assert.Equal(t, FloatSliceEquals(test.v1, test.v2), test.equals,
			"Test case: %+v", test)
	}
}
