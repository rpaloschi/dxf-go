package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type VertexTestSuite struct {
	suite.Suite
}

func (suite *VertexTestSuite) TestMinimalVertex() {
	expected := Vertex{
		BaseEntity: BaseEntity{
			Handle:    "LH",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		Location: core.Point{X: 1.1, Y: 1.2, Z: 1.3},
	}

	next := core.Tagger(strings.NewReader(testMinimalVertex))
	arc, err := NewVertex(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *VertexTestSuite) TestVertexAllAttribs() {
	expected := Vertex{
		BaseEntity: BaseEntity{
			Handle:        "ALL_ARGS",
			Owner:         "hb",
			Space:         PAPER,
			LayoutTabName: "layout",
			LayerName:     "L1",
			LineTypeName:  "CONTINUOUS",
			On:            true,
			Color:         2,
			LineWeight:    3,
			LineTypeScale: 2.5,
			Visible:       false,
			TrueColor:     core.TrueColor(0x684e45),
			ColorName:     "BROWN",
			Transparency:  5,
			ShadowMode:    CASTS,
		},
		Location:      core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		StartingWidth: 10.5,
		EndWidth:      15.8,
		Bulge:         11.2,
		Id:            3,
		CreatedByCurveFitting:    true,
		CurveFitTangentDefined:   true,
		SplineVertex:             true,
		SplineFrameCtrlPoint:     true,
		Is3dPolylineVertex:       true,
		Is3dPolylineMesh:         true,
		IsPolyfaceMeshVertex:     true,
		CurveFitTangentDirection: 0.2,
	}

	next := core.Tagger(strings.NewReader(testVertexAllAttribs))
	arc, err := NewVertex(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *VertexTestSuite) TestVertexNotEqualToDifferentType() {
	suite.False(Vertex{}.Equals(core.NewIntegerValue(0)))
}

func TestVertexTestSuite(t *testing.T) {
	suite.Run(t, new(VertexTestSuite))
}

const testMinimalVertex = `  0
VERTEX
  5
LH
  8
0
 10
1.1
 20
1.2
 30
1.3
`

const testVertexAllAttribs = `  0
VERTEX
  5
ALL_ARGS
  8
L1
  6
CONTINUOUS
 48
2.5
 60
1
 62
2
 67
1
284
1
330
hb
370
3
410
layout
420
6835781
430
BROWN
440
5
 10
1.1
 20
1.2
 30
1.3
 40
10.5
 41
15.8
 42
11.2
 91
3
 70
251
 50
0.2
`
