package core

import "math"

// MinFloatDelta value used to consider a floating point equal to 0, mainly to be used
// in internal equality functions.
// TODO: rpaloschi - consider the unit contained on the DXF itself.
const MinFloatDelta = 0.0000000001

// FloatEquals compares two floats using the MinFloatDelta difference to consider them
// Equals or not.
func FloatEquals(a float64, b float64) bool {
	return math.Abs(a-b) < MinFloatDelta
}

// FloatSliceEquals compares two slices of floats for equality.
// The slices are considered equals if both contains the same number of
// elements and FloatEquals returns true for all pairs of floats at the
// same index.
func FloatSliceEquals(a []float64, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, aValue := range a {
		bValue := b[i]

		if !FloatEquals(aValue, bValue) {
			return false
		}
	}

	return true
}
