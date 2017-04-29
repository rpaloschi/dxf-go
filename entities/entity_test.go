package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntitySliceEquality(t *testing.T) {
	testCases := []struct {
		e1     EntitySlice
		e2     EntitySlice
		equals bool
	}{
		{
			EntitySlice{},
			EntitySlice{},
			true,
		},
		{
			EntitySlice{&Line{}},
			EntitySlice{&Line{}},
			true,
		},
		{
			EntitySlice{&Vertex{Id: 1}, &Vertex{Id: 2}},
			EntitySlice{&Vertex{Id: 1}, &Vertex{Id: 2}},
			true,
		},
		{
			EntitySlice{&Vertex{Id: 1}},
			EntitySlice{},
			false,
		},
		{
			EntitySlice{&Vertex{Id: 1}, &Vertex{Id: 2}},
			EntitySlice{&Vertex{Id: 2}, &Vertex{Id: 1}},
			false,
		},
	}

	for i, test := range testCases {
		assert.Equal(t, test.equals, test.e1.Equals(test.e2), "Test index %v", i)
	}
}
