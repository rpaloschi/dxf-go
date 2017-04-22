package entities

import "github.com/rpaloschi/dxf-go/core"

// Circle Entity representation
type Circle struct {
	BaseEntity
	Thickness          float64
	Center             core.Point
	Radius             float64
	ExtrusionDirection core.Point
}

// Equals tests equality against another Circle.
func (c Circle) Equals(other core.DxfElement) bool {
	if otherCircle, ok := other.(*Circle); ok {
		return c.BaseEntity.Equals(otherCircle.BaseEntity) &&
			core.FloatEquals(c.Thickness, otherCircle.Thickness) &&
			c.Center.Equals(otherCircle.Center) &&
			core.FloatEquals(c.Radius, otherCircle.Radius) &&
			c.ExtrusionDirection.Equals(otherCircle.ExtrusionDirection)
	}
	return false
}

// NewCircle builds a new Circle from a slice of Tags.
func NewCircle(tags core.TagSlice) (*Circle, error) {
	circle := new(Circle)

	// set defaults
	circle.Radius = 1.0
	circle.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	circle.InitBaseEntityParser()
	circle.Update(map[int]core.TypeParser{
		39:  core.NewFloatTypeParserToVar(&circle.Thickness),
		10:  core.NewFloatTypeParserToVar(&circle.Center.X),
		20:  core.NewFloatTypeParserToVar(&circle.Center.Y),
		30:  core.NewFloatTypeParserToVar(&circle.Center.Z),
		40:  core.NewFloatTypeParserToVar(&circle.Radius),
		210: core.NewFloatTypeParserToVar(&circle.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&circle.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&circle.ExtrusionDirection.Z),
	})

	err := circle.Parse(tags)
	return circle, err
}
