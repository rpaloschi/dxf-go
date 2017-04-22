package entities

import "github.com/rpaloschi/dxf-go/core"

// Point Entity representation
type Point struct {
	BaseEntity
	Location           core.Point
	Thickness          float64
	ExtrusionDirection core.Point
	XAxisAngle         float64
}

// Equals tests equality against another Point.
func (c Point) Equals(other core.DxfElement) bool {
	if otherPoint, ok := other.(*Point); ok {
		return c.BaseEntity.Equals(otherPoint.BaseEntity) &&
			c.Location.Equals(otherPoint.Location) &&
			core.FloatEquals(c.Thickness, otherPoint.Thickness) &&
			c.ExtrusionDirection.Equals(otherPoint.ExtrusionDirection) &&
			core.FloatEquals(c.XAxisAngle, otherPoint.XAxisAngle)
	}
	return false
}

// NewPoint builds a new Point from a slice of Tags.
func NewPoint(tags core.TagSlice) (*Point, error) {
	point := new(Point)

	// set defaults
	point.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	point.InitBaseEntityParser()
	point.Update(map[int]core.TypeParser{

		10:  core.NewFloatTypeParserToVar(&point.Location.X),
		20:  core.NewFloatTypeParserToVar(&point.Location.Y),
		30:  core.NewFloatTypeParserToVar(&point.Location.Z),
		39:  core.NewFloatTypeParserToVar(&point.Thickness),
		50:  core.NewFloatTypeParserToVar(&point.XAxisAngle),
		210: core.NewFloatTypeParserToVar(&point.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&point.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&point.ExtrusionDirection.Z),
	})

	err := point.Parse(tags)
	return point, err
}
