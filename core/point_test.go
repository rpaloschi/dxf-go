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

func TestPointSliceEquality(t *testing.T) {
	testCases := []struct {
		p1     PointSlice
		p2     PointSlice
		equals bool
	}{
		{PointSlice{}, PointSlice{}, true},
		{PointSlice{Point{1.1, 2.2, 3.3}}, PointSlice{}, false},
		{PointSlice{}, PointSlice{Point{1.1, 2.2, 3.3}}, false},
		{
			PointSlice{Point{1.1, 2.2, 3.3}, Point{1.1, 2.2, 3.3}},
			PointSlice{Point{18.1, 82.2, 83.3}, Point{1.1, 2.2, 3.3}},
			false,
		},
		{
			PointSlice{Point{1.1, 2.2, 3.3}, Point{1.1, 2.2, 3.3}},
			PointSlice{Point{1.1, 2.2, 3.3}, Point{18.1, 82.2, 83.3}},
			false,
		},
		{PointSlice{Point{1.1, 2.2, 3.3}}, PointSlice{Point{1.1, 2.2, 3.3}}, true},
		{
			PointSlice{Point{1.1, 2.2, 3.3}, Point{18.1, 82.2, 83.3}},
			PointSlice{Point{1.1, 2.2, 3.3}, Point{18.1, 82.2, 83.3}},
			true,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.p1.Equals(test.p2), test.equals)
	}
}
