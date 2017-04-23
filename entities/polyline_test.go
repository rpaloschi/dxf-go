package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type PolylineTestSuite struct {
	suite.Suite
}

func (suite *PolylineTestSuite) TestMinimalPolyline() {
	expected := Polyline{
		BaseEntity: BaseEntity{
			Handle:    "3E5",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		Vertices:           VertexSlice{},
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalPolyline))
	arc, err := NewPolyline(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *PolylineTestSuite) TestPolylineAllAttribs() {
	expected := Polyline{
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
			ShadowMode:    CASTS_AND_RECEIVE,
		},
		Elevation:              1.5,
		Thickness:              3.3,
		Closed:                 true,
		CurveFitVerticesAdded:  true,
		SplineFitVerticesAdded: true,
		Is3dPolyline:           true,
		Is3dPolygonMesh:        true,
		PolygonMeshClosedNDir:  true,
		IsPolyfaceMesh:         true,
		LineTypeParentAround:   true,
		DefaultStartWidth:      0.5,
		DefaultEndWidth:        1.5,
		VertexCountM:           3,
		VertexCountN:           4,
		SmoothDensityM:         5,
		SmoothDensityN:         6,
		SmoothSurface:          CUBIC_BSPLINE,
		Vertices:               VertexSlice{},
		ExtrusionDirection:     core.Point{X: 32.1, Y: 12.6, Z: 95.1},
	}

	next := core.Tagger(strings.NewReader(testPolylineAllAttribs))
	arc, err := NewPolyline(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *PolylineTestSuite) TestPolylineNotEqualToDifferentType() {
	suite.False(Polyline{}.Equals(core.NewIntegerValue(0)))
}

func TestPolylineTestSuite(t *testing.T) {
	suite.Run(t, new(PolylineTestSuite))
}

const testMinimalPolyline = `  0
POLYLINE
  5
3E5
  8
0
`

const testPolylineAllAttribs = `  0
POLYLINE
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
0
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
 39
3.3
 30
1.5
 40
0.5
 41
1.5
 70
255
 71
3
 72
4
 73
5
 74
6
 75
6
210
32.1
220
12.6
230
95.1
`
