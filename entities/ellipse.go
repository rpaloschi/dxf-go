package entities

import "github.com/rpaloschi/dxf-go/core"

// Ellipse Entity representation
type Ellipse struct {
	BaseEntity
	Center                core.Point
	MajorAxisEnd          core.Point
	ExtrusionDirection    core.Point
	MinorToMajorAxisRatio float64
	StartParameter        float64
	EndParameter          float64
}

// Equals tests equality against another Ellipse.
func (e Ellipse) Equals(other core.DxfElement) bool {
	if otherEllipse, ok := other.(*Ellipse); ok {
		return e.BaseEntity.Equals(otherEllipse.BaseEntity) &&
			e.Center.Equals(otherEllipse.Center) &&
			e.MajorAxisEnd.Equals(otherEllipse.MajorAxisEnd) &&
			e.ExtrusionDirection.Equals(otherEllipse.ExtrusionDirection) &&
			core.FloatEquals(e.MinorToMajorAxisRatio, otherEllipse.MinorToMajorAxisRatio) &&
			core.FloatEquals(e.StartParameter, otherEllipse.StartParameter) &&
			core.FloatEquals(e.EndParameter, otherEllipse.EndParameter)
	}
	return false
}

// NewEllipse builds a new Ellipse from a slice of Tags.
func NewEllipse(tags core.TagSlice) (*Ellipse, error) {
	ellipse := new(Ellipse)

	// set default
	ellipse.MinorToMajorAxisRatio = 1.0
	ellipse.StartParameter = 0.0
	ellipse.EndParameter = 360.0
	ellipse.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	ellipse.InitBaseEntityParser()
	ellipse.Update(map[int]core.TypeParser{
		10:  core.NewFloatTypeParserToVar(&ellipse.Center.X),
		20:  core.NewFloatTypeParserToVar(&ellipse.Center.Y),
		30:  core.NewFloatTypeParserToVar(&ellipse.Center.Z),
		11:  core.NewFloatTypeParserToVar(&ellipse.MajorAxisEnd.X),
		21:  core.NewFloatTypeParserToVar(&ellipse.MajorAxisEnd.Y),
		31:  core.NewFloatTypeParserToVar(&ellipse.MajorAxisEnd.Z),
		210: core.NewFloatTypeParserToVar(&ellipse.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&ellipse.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&ellipse.ExtrusionDirection.Z),
		40:  core.NewFloatTypeParserToVar(&ellipse.MinorToMajorAxisRatio),
		41:  core.NewFloatTypeParserToVar(&ellipse.StartParameter),
		42:  core.NewFloatTypeParserToVar(&ellipse.EndParameter),
	})

	err := ellipse.Parse(tags)
	return ellipse, err
}
