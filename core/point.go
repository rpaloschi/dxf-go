package core

// Point 3d point representation
type Point struct {
	X float64
	Y float64
	Z float64
}

// Equals compares the point to the other Point for equality.
func (p Point) Equals(other Point) bool {
	return FloatEquals(p.X, other.X) &&
		FloatEquals(p.Y, other.Y) &&
		FloatEquals(p.Z, other.Z)
}

// PointSlice A slice of core.Point objects.
type PointSlice []Point

// Equals compares with other PointSlice for equality.
func (p PointSlice) Equals(other PointSlice) bool {
	if len(p) != len(other) {
		return false
	}

	for i, point1 := range p {
		point2 := other[i]

		if !point1.Equals(point2) {
			return false
		}
	}

	return true
}
