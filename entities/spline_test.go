package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type SplineTestSuite struct {
	suite.Suite
}

func (suite *SplineTestSuite) TestMinimalSpline() {
	expected := Spline{
		BaseEntity: BaseEntity{
			Handle:    "3E5",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		KnotValues:            []float64{},
		Weights:               []float64{},
		ControlPoints:         core.PointSlice{},
		FitPoints:             core.PointSlice{},
		KnotTolerance:         0.0000001,
		ControlPointTolerance: 0.0000001,
		FitTolerance:          0.0000000001,
	}

	next := core.Tagger(strings.NewReader(testMinimalSpline))
	spline, err := NewSpline(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(spline))

	suite.False(spline.IsSeqEnd())
	suite.False(spline.HasNestedEntities())
}

func (suite *SplineTestSuite) TestSplineAllAttribs() {
	expected := Spline{
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
			ShadowMode:    IGNORES,
		},
		NormalVector:          core.Point{X: 32.1, Y: 12.6, Z: 95.1},
		Closed:                true,
		Periodic:              true,
		Rational:              true,
		Planar:                true,
		Linear:                true,
		Degree:                3,
		KnotTolerance:         0.0000003,
		ControlPointTolerance: 0.0000004,
		FitTolerance:          0.0000000005,
		StartTangent:          core.Point{X: 1.6, Y: 2.5, Z: 3.4},
		EndTangent:            core.Point{X: 4.3, Y: 5.2, Z: 6.1},
		KnotValues:            []float64{3.1, 2.5, 1.6},
		Weights:               []float64{0.1, 0.5, 0.6},
		ControlPoints: core.PointSlice{
			core.Point{X: 1.0, Y: 2.0, Z: 3.0},
			core.Point{X: 2.6, Y: 2.1, Z: 3.2},
			core.Point{X: 3.1, Y: 2.8, Z: 3.8},
		},
		FitPoints: core.PointSlice{
			core.Point{X: 1.1, Y: 2.1, Z: 3.1},
			core.Point{X: 2.7, Y: 2.2, Z: 3.3},
			core.Point{X: 3.2, Y: 2.9, Z: 3.9},
		},
	}

	next := core.Tagger(strings.NewReader(testSplineAllAttribs))
	spline, err := NewSpline(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(spline))
}

func (suite *SplineTestSuite) TestSplineNotEqualToDifferentType() {
	suite.False(Spline{}.Equals(core.NewStringValue("STR")))
}

func TestSplineTestSuite(t *testing.T) {
	suite.Run(t, new(SplineTestSuite))
}

const testMinimalSpline = `  0
SPLINE
  5
3E5
  8
0
`

const testSplineAllAttribs = `  0
SPLINE
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
3
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
210
32.1
220
12.6
230
95.1
 70
31
 71
3
 42
0.0000003
 43
0.0000004
 44
0.0000000005
 12
1.6
 22
2.5
 32
3.4
 13
4.3
 23
5.2
 33
6.1
 40
3.1
 40
2.5
 40
1.6
 41
0.1
 41
0.5
 41
0.6
 10
1.0
 20
2.0
 30
3.0
 10
2.6
 20
2.1
 30
3.2
 10
3.1
 20
2.8
 30
3.8
 11
1.1
 21
2.1
 31
3.1
 11
2.7
 21
2.2
 31
3.3
 11
3.2
 21
2.9
 31
3.9
`
