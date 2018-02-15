package entities

import "github.com/rpaloschi/dxf-go/core"

// Vertex Entity representation
type Vertex struct {
	BaseEntity
	Location                 core.Point
	StartingWidth            float64
	EndWidth                 float64
	Bulge                    float64
	CreatedByCurveFitting    bool
	CurveFitTangentDefined   bool
	SplineVertex             bool
	SplineFrameCtrlPoint     bool
	Is3dPolylineVertex       bool
	Is3dPolylineMesh         bool
	IsPolyfaceMeshVertex     bool
	CurveFitTangentDirection float64
	Id                       int64
}

// Equals tests equality against another Vertex.
func (c Vertex) Equals(other core.DxfElement) bool {
	if otherVertex, ok := other.(*Vertex); ok {
		return c.BaseEntity.Equals(otherVertex.BaseEntity) &&
			c.Location.Equals(otherVertex.Location) &&
			core.FloatEquals(c.StartingWidth, otherVertex.StartingWidth) &&
			core.FloatEquals(c.EndWidth, otherVertex.EndWidth) &&
			core.FloatEquals(c.Bulge, otherVertex.Bulge) &&
			c.CreatedByCurveFitting == otherVertex.CreatedByCurveFitting &&
			c.CurveFitTangentDefined == otherVertex.CurveFitTangentDefined &&
			c.SplineVertex == otherVertex.SplineVertex &&
			c.SplineFrameCtrlPoint == otherVertex.SplineFrameCtrlPoint &&
			c.Is3dPolylineVertex == otherVertex.Is3dPolylineVertex &&
			c.Is3dPolylineMesh == otherVertex.Is3dPolylineMesh &&
			c.IsPolyfaceMeshVertex == otherVertex.IsPolyfaceMeshVertex &&
			core.FloatEquals(c.CurveFitTangentDirection,
				otherVertex.CurveFitTangentDirection) &&
			c.Id == otherVertex.Id
	}
	return false
}

const extraVertexCurveFittingBit = 0x1
const curveFitTangentDefinedBit = 0x2
const splineVertexCreatedBit = 0x8
const splineFrameCtrlPointBit = 0x10
const polylineVertex3dBit = 0x20
const polygonMesh3dBit = 0x40
const polyfaceMeshVertexBit = 0x80

// NewVertex builds a new Vertex from a slice of Tags.
func NewVertex(tags core.TagSlice) (*Vertex, error) {
	vertex := new(Vertex)

	vertex.InitBaseEntityParser()
	vertex.Update(map[int]core.TypeParser{
		10: core.NewFloatTypeParserToVar(&vertex.Location.X),
		20: core.NewFloatTypeParserToVar(&vertex.Location.Y),
		30: core.NewFloatTypeParserToVar(&vertex.Location.Z),
		40: core.NewFloatTypeParserToVar(&vertex.StartingWidth),
		41: core.NewFloatTypeParserToVar(&vertex.EndWidth),
		42: core.NewFloatTypeParserToVar(&vertex.Bulge),
		50: core.NewFloatTypeParserToVar(&vertex.CurveFitTangentDirection),
		70: core.NewIntTypeParser(func(flags int64) {
			vertex.CreatedByCurveFitting = flags&extraVertexCurveFittingBit != 0
			vertex.CurveFitTangentDefined = flags&curveFitTangentDefinedBit != 0
			vertex.SplineVertex = flags&splineVertexCreatedBit != 0
			vertex.SplineFrameCtrlPoint = flags&splineFrameCtrlPointBit != 0
			vertex.Is3dPolylineVertex = flags&polylineVertex3dBit != 0
			vertex.Is3dPolylineMesh = flags&polygonMesh3dBit != 0
			vertex.IsPolyfaceMeshVertex = flags&polyfaceMeshVertexBit != 0
		}),
		91: core.NewIntTypeParserToVar(&vertex.Id),
	})

	err := vertex.Parse(tags)
	return vertex, err
}

// VertexSlice a slice of Vertex objects.
type VertexSlice []*Vertex

// Equals Compares two Vertices for equality.
func (v VertexSlice) Equals(other VertexSlice) bool {
	if len(v) != len(other) {
		return false
	}

	for i, vertex := range v {
		otherVertex := other[i]

		if !vertex.Equals(otherVertex) {
			return false
		}
	}

	return true
}
