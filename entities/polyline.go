package entities

import (
	"github.com/rpaloschi/dxf-go/core"
)

// Polyline Entity representation
type Polyline struct {
	BaseEntity
	Elevation              float64
	Thickness              float64
	Closed                 bool
	CurveFitVerticesAdded  bool
	SplineFitVerticesAdded bool
	Is3dPolyline           bool
	Is3dPolygonMesh        bool
	PolygonMeshClosedNDir  bool
	IsPolyfaceMesh         bool
	LineTypeParentAround   bool
	DefaultStartWidth      float64
	DefaultEndWidth        float64
	VertexCountM           int64
	VertexCountN           int64
	SmoothDensityM         int64
	SmoothDensityN         int64
	SmoothSurface          SmoothSurfaceType
	ExtrusionDirection     core.Point
	Vertices               VertexSlice
}

// SmoothSurfaceType representation
type SmoothSurfaceType int

const (
	NO_SMOOTH_SURFACE_FITTED SmoothSurfaceType = 0
	QUADRATIC_BSPLINE        SmoothSurfaceType = 5
	CUBIC_BSPLINE            SmoothSurfaceType = 6
	BEZIER                   SmoothSurfaceType = 8
)

// Equals tests equality against another Polyline.
func (p Polyline) Equals(other core.DxfElement) bool {
	if otherPolyline, ok := other.(*Polyline); ok {
		return p.BaseEntity.Equals(otherPolyline.BaseEntity) &&
			core.FloatEquals(p.Elevation, otherPolyline.Elevation) &&
			core.FloatEquals(p.Thickness, otherPolyline.Thickness) &&
			p.Closed == otherPolyline.Closed &&
			p.CurveFitVerticesAdded == otherPolyline.CurveFitVerticesAdded &&
			p.SplineFitVerticesAdded == otherPolyline.SplineFitVerticesAdded &&
			p.Is3dPolyline == otherPolyline.Is3dPolyline &&
			p.Is3dPolygonMesh == otherPolyline.Is3dPolygonMesh &&
			p.PolygonMeshClosedNDir == otherPolyline.PolygonMeshClosedNDir &&
			p.IsPolyfaceMesh == otherPolyline.IsPolyfaceMesh &&
			p.LineTypeParentAround == otherPolyline.LineTypeParentAround &&
			core.FloatEquals(p.DefaultStartWidth, otherPolyline.DefaultStartWidth) &&
			core.FloatEquals(p.DefaultEndWidth, otherPolyline.DefaultEndWidth) &&
			p.VertexCountM == otherPolyline.VertexCountM &&
			p.VertexCountN == otherPolyline.VertexCountN &&
			p.SmoothDensityM == otherPolyline.SmoothDensityM &&
			p.SmoothDensityN == otherPolyline.SmoothDensityN &&
			p.SmoothSurface == otherPolyline.SmoothSurface &&
			p.ExtrusionDirection.Equals(otherPolyline.ExtrusionDirection) &&
			p.Vertices.Equals(otherPolyline.Vertices)
	}
	return false
}

// HasNestedEntities a Polyline will have nested entities.
func (p Polyline) HasNestedEntities() bool {
	return true
}

// AddNestedEntities a Polyline will contain only Vertex as nested entities,
// other types are simply ignored.
func (p *Polyline) AddNestedEntities(entities EntitySlice) {
	for _, entity := range entities {
		if vertex, ok := entity.(*Vertex); ok {
			p.Vertices = append(p.Vertices, vertex)
		} else {
			core.Log.Printf(
				"Skipping entity %v. Polylines can only contain Vertex entities.",
				entity)
		}
	}
}

const closedPolylineBit = 0x1
const curveFitVerticesAddedBit = 0x2
const splineFitVerticesAddedBit = 0x4
const is3dPolylineBit = 0x8
const is3dPolygonMeshBit = 0x10
const closedNDirectionBit = 0x20
const polyfaceMeshBit = 0x40
const lineTypePatternBit = 0x80

// NewPolyline builds a new Polyline from a slice of Tags (without vertices)
func NewPolyline(tags core.TagSlice) (*Polyline, error) {
	polyline := new(Polyline)

	// set defaults
	polyline.Vertices = make(VertexSlice, 0)
	polyline.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	polyline.InitBaseEntityParser()
	polyline.Update(map[int]core.TypeParser{
		30: core.NewFloatTypeParserToVar(&polyline.Elevation),
		39: core.NewFloatTypeParserToVar(&polyline.Thickness),
		40: core.NewFloatTypeParserToVar(&polyline.DefaultStartWidth),
		41: core.NewFloatTypeParserToVar(&polyline.DefaultEndWidth),
		70: core.NewIntTypeParser(func(flags int64) {
			polyline.Closed = flags&closedPolylineBit != 0
			polyline.CurveFitVerticesAdded = flags&curveFitVerticesAddedBit != 0
			polyline.SplineFitVerticesAdded = flags&splineFitVerticesAddedBit != 0
			polyline.Is3dPolyline = flags&is3dPolylineBit != 0
			polyline.Is3dPolygonMesh = flags&is3dPolygonMeshBit != 0
			polyline.PolygonMeshClosedNDir = flags&closedNDirectionBit != 0
			polyline.IsPolyfaceMesh = flags&polyfaceMeshBit != 0
			polyline.LineTypeParentAround = flags&lineTypePatternBit != 0
		}),
		71: core.NewIntTypeParserToVar(&polyline.VertexCountM),
		72: core.NewIntTypeParserToVar(&polyline.VertexCountN),
		73: core.NewIntTypeParserToVar(&polyline.SmoothDensityM),
		74: core.NewIntTypeParserToVar(&polyline.SmoothDensityN),
		75: core.NewIntTypeParser(func(value int64) {
			polyline.SmoothSurface = SmoothSurfaceType(value)
		}),
		210: core.NewFloatTypeParserToVar(&polyline.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&polyline.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&polyline.ExtrusionDirection.Z),
	})

	err := polyline.Parse(tags)
	return polyline, err
}
