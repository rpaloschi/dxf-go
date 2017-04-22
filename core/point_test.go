package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointEquality(t *testing.T) {
	testCases := []struct {
		p1     Point
		p2     Point
		equals bool
	}{
		{Point{1.1, 2.2, 3.3}, Point{1.1, 2.2, 3.3}, true},
		{Point{1.1, 2.2, 3.3}, Point{}, false},
	}

	for _, test := range testCases {
		assert.Equal(t, test.p1.Equals(test.p2), test.equals)
	}
}
