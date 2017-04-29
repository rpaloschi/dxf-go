package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type CircleTestSuite struct {
	suite.Suite
}

func (suite *CircleTestSuite) TestMinimalCircle() {
	expected := Circle{
		BaseEntity: BaseEntity{
			Handle:    "3E5",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		Center:             core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		Radius:             5.0,
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalCircle))
	circle, err := NewCircle(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(circle))

	suite.False(circle.IsSeqEnd())
	suite.False(circle.HasNestedEntities())
}

func (suite *CircleTestSuite) TestCircleAllAttribs() {
	expected := Circle{
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
		ExtrusionDirection: core.Point{X: 32.1, Y: 12.6, Z: 95.1},
	}

	next := core.Tagger(strings.NewReader(testCircleAllAttribs))
	circle, err := NewCircle(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(circle))
}

func (suite *CircleTestSuite) TestCircleOff() {
	expected := Circle{
		BaseEntity: BaseEntity{
			On:      false,
			Color:   4,
			Visible: true,
		},
		Center:             core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		Radius:             5.0,
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testCircleOff))
	circle, err := NewCircle(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(circle))
}

func (suite *CircleTestSuite) TestCircleNotEqualToDifferentType() {
	suite.False(Circle{}.Equals(core.NewIntegerValue(0)))
}

func TestCircleTestSuite(t *testing.T) {
	suite.Run(t, new(CircleTestSuite))
}

const testMinimalCircle = `  0
CIRCLE
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
`

const testCircleAllAttribs = `  0
CIRCLE
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
210
32.1
220
12.6
230
95.1
`

const testCircleOff = `  0
CIRCLE
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
`
