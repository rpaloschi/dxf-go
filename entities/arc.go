package entities

import "github.com/rpaloschi/dxf-go/core"

// Arc Entity representation
type Arc struct {
	BaseEntity
	Thickness          float64
	Center             core.Point
	Radius             float64
	StartAngle         float64
	EndAngle           float64
	ExtrusionDirection core.Point
}

// Equals tests equality against another Arc.
func (a Arc) Equals(other core.DxfElement) bool {
	if otherArc, ok := other.(*Arc); ok {
		return a.BaseEntity.Equals(otherArc.BaseEntity) &&
			core.FloatEquals(a.Thickness, otherArc.Thickness) &&
			a.Center.Equals(otherArc.Center) &&
			core.FloatEquals(a.Radius, otherArc.Radius) &&
			core.FloatEquals(a.StartAngle, otherArc.StartAngle) &&
			core.FloatEquals(a.EndAngle, otherArc.EndAngle) &&
			a.ExtrusionDirection.Equals(otherArc.ExtrusionDirection)
	}
	return false
}

// NewArc builds a new Arc from a slice of Tags.
func NewArc(tags core.TagSlice) (*Arc, error) {
	arc := new(Arc)

	// set defaults (the ones which the zero value is the same are not initialized here)
	arc.Radius = 1.0
	arc.EndAngle = 360.0
	arc.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	arc.InitBaseEntityParser()
	arc.Update(map[int]core.TypeParser{
		39:  core.NewFloatTypeParserToVar(&arc.Thickness),
		10:  core.NewFloatTypeParserToVar(&arc.Center.X),
		20:  core.NewFloatTypeParserToVar(&arc.Center.Y),
		30:  core.NewFloatTypeParserToVar(&arc.Center.Z),
		40:  core.NewFloatTypeParserToVar(&arc.Radius),
		50:  core.NewFloatTypeParserToVar(&arc.StartAngle),
		51:  core.NewFloatTypeParserToVar(&arc.EndAngle),
		210: core.NewFloatTypeParserToVar(&arc.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&arc.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&arc.ExtrusionDirection.Z),
	})

	err := arc.Parse(tags)
	return arc, err
}
