package entities

import "github.com/rpaloschi/dxf-go/core"

// Spline Entity representation
type Spline struct {
	BaseEntity
	NormalVector          core.Point
	Closed                bool
	Periodic              bool
	Rational              bool
	Planar                bool
	Linear                bool
	Degree                int
	KnotTolerance         float64
	ControlPointTolerance float64
	FitTolerance          float64
	StartTangent          core.Point
	EndTangent            core.Point
	KnotValues            []float64
	Weights               []float64
	ControlPoints         core.PointSlice
	FitPoints             core.PointSlice
}

// Equals tests equality against another Spline.
func (s Spline) Equals(other core.DxfElement) bool {
	if otherSpline, ok := other.(*Spline); ok {
		return s.BaseEntity.Equals(otherSpline.BaseEntity) &&
			s.NormalVector.Equals(otherSpline.NormalVector) &&
			s.Closed == otherSpline.Closed &&
			s.Periodic == otherSpline.Periodic &&
			s.Rational == otherSpline.Rational &&
			s.Planar == otherSpline.Planar &&
			s.Linear == otherSpline.Linear &&
			s.Degree == otherSpline.Degree &&
			core.FloatEquals(s.KnotTolerance, otherSpline.KnotTolerance) &&
			core.FloatEquals(s.ControlPointTolerance,
				otherSpline.ControlPointTolerance) &&
			core.FloatEquals(s.FitTolerance, otherSpline.FitTolerance) &&
			s.StartTangent.Equals(otherSpline.StartTangent) &&
			s.EndTangent.Equals(otherSpline.EndTangent) &&
			core.FloatSliceEquals(s.KnotValues, otherSpline.KnotValues) &&
			core.FloatSliceEquals(s.Weights, otherSpline.Weights) &&
			s.ControlPoints.Equals(otherSpline.ControlPoints) &&
			s.FitPoints.Equals(otherSpline.FitPoints)
	}
	return false
}

const closedSplineBit = 0x1
const periodicSplineBit = 0x2
const rationalSplineBit = 0x4
const planarBit = 0x8
const linearBit = 0x10

// NewSpline builds a new Spline from a slice of Tags.
func NewSpline(tags core.TagSlice) (*Spline, error) {
	spline := new(Spline)

	// set defaults
	spline.KnotTolerance = 0.0000001
	spline.ControlPointTolerance = 0.0000001
	spline.FitTolerance = 0.0000000001
	spline.KnotValues = make([]float64, 0)
	spline.Weights = make([]float64, 0)
	spline.ControlPoints = make(core.PointSlice, 0)
	spline.FitPoints = make(core.PointSlice, 0)

	spline.InitBaseEntityParser()
	spline.Update(map[int]core.TypeParser{
		10: core.NewFloatTypeParser(func(value float64) {
			spline.ControlPoints = append(spline.ControlPoints,
				core.Point{X: value})
		}),
		20: core.NewFloatTypeParser(func(value float64) {
			spline.ControlPoints[len(spline.ControlPoints)-1].Y = value
		}),
		30: core.NewFloatTypeParser(func(value float64) {
			spline.ControlPoints[len(spline.ControlPoints)-1].Z = value
		}),
		11: core.NewFloatTypeParser(func(value float64) {
			spline.FitPoints = append(spline.FitPoints,
				core.Point{X: value})
		}),
		21: core.NewFloatTypeParser(func(value float64) {
			spline.FitPoints[len(spline.FitPoints)-1].Y = value
		}),
		31: core.NewFloatTypeParser(func(value float64) {
			spline.FitPoints[len(spline.FitPoints)-1].Z = value
		}),
		12: core.NewFloatTypeParserToVar(&spline.StartTangent.X),
		22: core.NewFloatTypeParserToVar(&spline.StartTangent.Y),
		32: core.NewFloatTypeParserToVar(&spline.StartTangent.Z),
		13: core.NewFloatTypeParserToVar(&spline.EndTangent.X),
		23: core.NewFloatTypeParserToVar(&spline.EndTangent.Y),
		33: core.NewFloatTypeParserToVar(&spline.EndTangent.Z),
		40: core.NewFloatTypeParser(func(value float64) {
			spline.KnotValues = append(spline.KnotValues, value)
		}),
		41: core.NewFloatTypeParser(func(value float64) {
			spline.Weights = append(spline.Weights, value)
		}),
		42: core.NewFloatTypeParserToVar(&spline.KnotTolerance),
		43: core.NewFloatTypeParserToVar(&spline.ControlPointTolerance),
		44: core.NewFloatTypeParserToVar(&spline.FitTolerance),
		70: core.NewIntTypeParser(func(flags int) {
			spline.Closed = flags&closedSplineBit != 0
			spline.Periodic = flags&periodicSplineBit != 0
			spline.Rational = flags&rationalSplineBit != 0
			spline.Planar = flags&planarBit != 0
			spline.Linear = flags&linearBit != 0
		}),
		71: core.NewIntTypeParserToVar(&spline.Degree),
		// We don't need those as we accumulate the values.
		// leaving them here as documentation and to remember to generate
		// them on writing.
		// 72:  core.NewIntTypeParserToVar(&spline.NrKnots),
		// 73:  core.NewIntTypeParserToVar(&spline.NrControlPoints),
		// 74:  core.NewIntTypeParserToVar(&spline.NrFitPoints),
		210: core.NewFloatTypeParserToVar(&spline.NormalVector.X),
		220: core.NewFloatTypeParserToVar(&spline.NormalVector.Y),
		230: core.NewFloatTypeParserToVar(&spline.NormalVector.Z),
	})

	err := spline.Parse(tags)
	return spline, err
}
