package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type LWPolylineTestSuite struct {
	suite.Suite
}

func (suite *LWPolylineTestSuite) TestMinimalLWPolyline() {
	expected := LWPolyline{
		BaseEntity: BaseEntity{
			Handle:    "LWP",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalLWPolyline))
	polyline, err := NewLWPolyline(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(polyline))

	suite.False(polyline.IsSeqEnd())
	suite.False(polyline.HasNestedEntities())
}

func (suite *LWPolylineTestSuite) TestLWPolylineAllAttribs() {
	expected := LWPolyline{
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
			ShadowMode:    RECEIVES,
		},
		Closed:        true,
		Plinegen:      true,
		ConstantWidth: 3.7,
		Elevation:     12.3,
		Thickness:     9.1,
		Points: LWPolyLinePointSlice{
			LWPolyLinePoint{
				Point:         core.Point{X: 1, Y: 2},
				Id:            0,
				StartingWidth: 1.0,
				EndWidth:      2.0,
				Bulge:         0.5,
			},
			LWPolyLinePoint{
				Point:         core.Point{X: 5, Y: 6.1},
				Id:            1,
				StartingWidth: 3.0,
				EndWidth:      2.5,
				Bulge:         5.5,
			},
			LWPolyLinePoint{
				Point:         core.Point{X: 1, Y: 1},
				Id:            2,
				StartingWidth: 1.0,
				EndWidth:      2.0,
				Bulge:         4.5,
			},
		},
		ExtrusionDirection: core.Point{X: 32.1, Y: 12.6, Z: 95.1},
	}

	next := core.Tagger(strings.NewReader(testLWPolylineAllAttribs))
	polyline, err := NewLWPolyline(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(polyline))
}

func (suite *LWPolylineTestSuite) TestLWPolylineNotEqualToDifferentType() {
	suite.False(LWPolyline{}.Equals(core.NewFloatValue(0.1)))
}

func TestLWPolylineTestSuite(t *testing.T) {
	suite.Run(t, new(LWPolylineTestSuite))
}

var pPoint1 = LWPolyLinePoint{
	Point:         core.Point{X: 1, Y: 2, Z: 3},
	Id:            0,
	StartingWidth: 1.0,
	EndWidth:      2.0,
	Bulge:         0.5,
}

var pPoint2 = LWPolyLinePoint{
	Point:         core.Point{X: 1, Y: 2, Z: 3},
	Id:            0,
	StartingWidth: 2.0,
	EndWidth:      3.0,
	Bulge:         1.5,
}

func TestPolyLinePointEquality(t *testing.T) {
	testCases := []struct {
		p1     LWPolyLinePoint
		p2     LWPolyLinePoint
		equals bool
	}{
		{
			pPoint1,
			pPoint1,
			true,
		},
		{
			pPoint1,
			LWPolyLinePoint{
				Point:         core.Point{X: 1, Y: 2, Z: 3},
				Id:            0,
				StartingWidth: 1.0,
				EndWidth:      2.0,
				Bulge:         0.5,
			},
			true,
		},
		{
			pPoint1,
			LWPolyLinePoint{
				Point:         core.Point{X: 1, Y: 3, Z: 3},
				Id:            0,
				StartingWidth: 1.0,
				EndWidth:      2.0,
				Bulge:         0.5,
			},
			false,
		},
		{
			pPoint1,
			LWPolyLinePoint{
				Point:         core.Point{X: 1, Y: 2, Z: 3},
				Id:            1,
				StartingWidth: 1.0,
				EndWidth:      2.0,
				Bulge:         0.5,
			},
			false,
		},
		{
			pPoint1,
			pPoint2,
			false,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.equals, test.p1.Equals(test.p2))
	}
}

func TestPolyLinePointSliceEquality(t *testing.T) {
	testCases := []struct {
		p1     LWPolyLinePointSlice
		p2     LWPolyLinePointSlice
		equals bool
	}{
		{
			LWPolyLinePointSlice{},
			LWPolyLinePointSlice{},
			true,
		},
		{
			LWPolyLinePointSlice{pPoint1},
			LWPolyLinePointSlice{pPoint1},
			true,
		},
		{
			LWPolyLinePointSlice{pPoint1, pPoint2},
			LWPolyLinePointSlice{pPoint1, pPoint2},
			true,
		},
		{
			LWPolyLinePointSlice{pPoint1},
			LWPolyLinePointSlice{},
			false,
		},
		{
			LWPolyLinePointSlice{pPoint1, pPoint2},
			LWPolyLinePointSlice{pPoint2, pPoint1},
			false,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.equals, test.p1.Equals(test.p2))
	}
}

const testMinimalLWPolyline = `  0
LWPOLYLINE
  5
LWP
  8
0
`

const testLWPolylineAllAttribs = `  0
LWPOLYLINE
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
2
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
 90
3
 70
129
 38
12.3
 39
9.1
 43
3.7
210
32.1
220
12.6
230
95.1
 10
1.0
 20
2.0
 91
0
 40
1.0
 41
2.0
 42
0.5
 10
5.0
 20
6.1
 91
1
 40
3.0
 41
2.5
 42
5.5
 10
1.0
 20
1.0
 91
2
 40
1.0
 41
2.0
 42
4.5
`
