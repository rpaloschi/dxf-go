package core

// Point 3d point representation
type Point struct {
	X float64
	Y float64
	Z float64
}

// Equals compare it to the other Point for equality.
func (p Point) Equals(other Point) bool {
	return FloatEquals(p.X, other.X) &&
		FloatEquals(p.Y, other.Y) &&
		FloatEquals(p.Z, other.Z)
}
