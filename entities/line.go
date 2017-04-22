package entities

import "github.com/rpaloschi/dxf-go/core"

// Line Entity representation
type Line struct {
	BaseEntity
	Thickness          float64
	Start              core.Point
	End                core.Point
	ExtrusionDirection core.Point
}

// Equals tests equality against another Line.
func (a Line) Equals(other core.DxfElement) bool {
	if otherLine, ok := other.(*Line); ok {
		return a.BaseEntity.Equals(otherLine.BaseEntity) &&
			core.FloatEquals(a.Thickness, otherLine.Thickness) &&
			a.Start.Equals(otherLine.Start) &&
			a.End.Equals(otherLine.End) &&
			a.ExtrusionDirection.Equals(otherLine.ExtrusionDirection)
	}
	return false
}

// NewLine builds a new Line from a slice of Tags.
func NewLine(tags core.TagSlice) (*Line, error) {
	line := new(Line)

	// set defaults
	line.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	line.InitBaseEntityParser()
	line.Update(map[int]core.TypeParser{
		39:  core.NewFloatTypeParserToVar(&line.Thickness),
		10:  core.NewFloatTypeParserToVar(&line.Start.X),
		20:  core.NewFloatTypeParserToVar(&line.Start.Y),
		30:  core.NewFloatTypeParserToVar(&line.Start.Z),
		11:  core.NewFloatTypeParserToVar(&line.End.X),
		21:  core.NewFloatTypeParserToVar(&line.End.Y),
		31:  core.NewFloatTypeParserToVar(&line.End.Z),
		210: core.NewFloatTypeParserToVar(&line.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&line.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&line.ExtrusionDirection.Z),
	})

	err := line.Parse(tags)
	return line, err
}
