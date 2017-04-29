package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type ArcTestSuite struct {
	suite.Suite
}

func (suite *ArcTestSuite) TestMinimalArc() {
	expected := Arc{
		BaseEntity: BaseEntity{
			Handle:    "3E5",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		Center:             core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		Radius:             5.0,
		StartAngle:         10.0,
		EndAngle:           90.0,
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalArc))
	arc, err := NewArc(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))

	suite.False(arc.IsSeqEnd())
	suite.False(arc.HasNestedEntities())
}

func (suite *ArcTestSuite) TestArcAllAttribs() {
	expected := Arc{
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
		Thickness:          3.3,
		Center:             core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		Radius:             5.0,
		StartAngle:         15.0,
		EndAngle:           45.0,
		ExtrusionDirection: core.Point{X: 32.1, Y: 12.6, Z: 95.1},
	}

	next := core.Tagger(strings.NewReader(testArcAllAttribs))
	arc, err := NewArc(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *ArcTestSuite) TestArcOff() {
	expected := Arc{
		BaseEntity: BaseEntity{
			On:      false,
			Color:   4,
			Visible: true,
		},
		Center:             core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		Radius:             5.0,
		StartAngle:         15.0,
		EndAngle:           45.0,
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testArcOff))
	arc, err := NewArc(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *ArcTestSuite) TestArcNotEqualToDifferentType() {
	suite.False(Arc{}.Equals(core.NewIntegerValue(0)))
}

func TestArcTestSuite(t *testing.T) {
	suite.Run(t, new(ArcTestSuite))
}

const testMinimalArc = `  0
ARC
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
 40
5.0
 50
10.0
 51
90.0
`

const testArcAllAttribs = `  0
ARC
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
 39
3.3
 10
1.1
 20
1.2
 30
1.3
 40
5.0
 50
15.0
 51
45.0
210
32.1
220
12.6
230
95.1
`

const testArcOff = `  0
ARC
 62
-4
 10
1.1
 20
1.2
 30
1.3
 40
5.0
 50
15.0
 51
45.0
`
