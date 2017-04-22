package entities

import "github.com/rpaloschi/dxf-go/core"

// LWPolyline Entity representation
type LWPolyline struct {
	BaseEntity
	Closed             bool
	Plinegen           bool
	ConstantWidth      float64
	Elevation          float64
	Thickness          float64
	Points             LWPolyLinePointSlice
	ExtrusionDirection core.Point
}

// Equals tests equality against another LWPolyline.
func (p LWPolyline) Equals(other core.DxfElement) bool {
	if otherLWPolyline, ok := other.(*LWPolyline); ok {
		return p.BaseEntity.Equals(otherLWPolyline.BaseEntity) &&
			p.Closed == otherLWPolyline.Closed &&
			p.Plinegen == otherLWPolyline.Plinegen &&
			core.FloatEquals(p.ConstantWidth, otherLWPolyline.ConstantWidth) &&
			core.FloatEquals(p.Elevation, otherLWPolyline.Elevation) &&
			core.FloatEquals(p.Thickness, otherLWPolyline.Thickness) &&
			p.Points.Equals(otherLWPolyline.Points) &&
			p.ExtrusionDirection.Equals(otherLWPolyline.ExtrusionDirection)
	}
	return false
}

const closedBit = 0x1
const plinegenBit = 0x80

// NewLWPolyline builds a new LWPolyline from a slice of Tags.
func NewLWPolyline(tags core.TagSlice) (*LWPolyline, error) {
	polyline := new(LWPolyline)

	// set defaults
	polyline.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	polyline.InitBaseEntityParser()

	pointIndex := -1
	polyline.Update(map[int]core.TypeParser{
		70: core.NewIntTypeParser(func(flags int) {
			polyline.Closed = flags&closedBit != 0
			polyline.Plinegen = flags&plinegenBit != 0
		}),
		90: core.NewIntTypeParser(func(value int) {
			polyline.Points = make(LWPolyLinePointSlice, value)
		}),
		38: core.NewFloatTypeParserToVar(&polyline.Elevation),
		39: core.NewFloatTypeParserToVar(&polyline.Thickness),
		43: core.NewFloatTypeParserToVar(&polyline.ConstantWidth),
		10: core.NewFloatTypeParser(func(x float64) {
			pointIndex++
			polyline.Points[pointIndex].Point.X = x
		}),
		20: core.NewFloatTypeParser(func(y float64) {
			polyline.Points[pointIndex].Point.Y = y
		}),
		91: core.NewIntTypeParser(func(value int) {
			polyline.Points[pointIndex].Id = value
		}),
		40: core.NewFloatTypeParser(func(value float64) {
			polyline.Points[pointIndex].StartingWidth = value
		}),
		41: core.NewFloatTypeParser(func(value float64) {
			polyline.Points[pointIndex].EndWidth = value
		}),
		42: core.NewFloatTypeParser(func(value float64) {
			polyline.Points[pointIndex].Bulge = value
		}),
		210: core.NewFloatTypeParserToVar(&polyline.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&polyline.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&polyline.ExtrusionDirection.Z),
	})

	err := polyline.Parse(tags)
	return polyline, err
}

// LWPolyLinePointSlice Slice of LWPolyLinePoint
type LWPolyLinePointSlice []LWPolyLinePoint

// Equals Compares two LWPolyLinePointSlices for equality.
func (p LWPolyLinePointSlice) Equals(other LWPolyLinePointSlice) bool {
	if len(p) != len(other) {
		return false
	}

	for i, point := range p {
		otherPoint := other[i]

		if !point.Equals(otherPoint) {
			return false
		}
	}

	return true
}

// LWPolyLinePoint point and attributes in an LWPolyline.
type LWPolyLinePoint struct {
	Point         core.Point
	Id            int
	StartingWidth float64
	EndWidth      float64
	Bulge         float64
}

// Equals compares two LWPolyLinePoints for equality
func (p LWPolyLinePoint) Equals(other LWPolyLinePoint) bool {
	return p.Point.Equals(other.Point) &&
		p.Id == other.Id &&
		core.FloatEquals(p.StartingWidth, other.StartingWidth) &&
		core.FloatEquals(p.EndWidth, other.EndWidth) &&
		core.FloatEquals(p.Bulge, other.Bulge)
}
