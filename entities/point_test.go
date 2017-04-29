package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type PointTestSuite struct {
	suite.Suite
}

func (suite *PointTestSuite) TestMinimalPoint() {
	expected := Point{
		BaseEntity: BaseEntity{
			Handle:    "3E5",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		Location:           core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalPoint))
	point, err := NewPoint(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(point))

	suite.False(point.IsSeqEnd())
	suite.False(point.HasNestedEntities())
}

func (suite *PointTestSuite) TestPointAllAttribs() {
	expected := Point{
		BaseEntity: BaseEntity{
			Handle:        "ALL_ARGS",
			Owner:         "hb",
			Space:         MODEL,
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
		Location:           core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		Thickness:          3.3,
		ExtrusionDirection: core.Point{X: 32.1, Y: 12.6, Z: 95.1},
		XAxisAngle:         5.0,
	}

	next := core.Tagger(strings.NewReader(testPointAllAttribs))
	point, err := NewPoint(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(point))
}

func (suite *PointTestSuite) TestPointNotEqualToDifferentType() {
	suite.False(Point{}.Equals(core.NewStringValue("AAA")))
}

func TestPointTestSuite(t *testing.T) {
	suite.Run(t, new(PointTestSuite))
}

const testMinimalPoint = `  0
POINT
  5
3E5
  8
0
 10
1.1
 20
1.2
 30
1.3
`

const testPointAllAttribs = `  0
POINT
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
0
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
 39
3.3
 10
1.1
 20
1.2
 30
1.3
 50
5.0
210
32.1
220
12.6
230
95.1
`
