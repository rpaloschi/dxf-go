package core

import "math"

// MinFloatDelta value used to consider a floating point equal to 0, mainly to be used
// in internal equality functions.
const MinFloatDelta = 0.001

// FloatEquals compares two floats using the MinFloatDelta difference to consider them
// Equals or not.
func FloatEquals(a float64, b float64) bool {
	return math.Abs(a-b) < MinFloatDelta
}
