package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type EllipseTestSuite struct {
	suite.Suite
}

func (suite *EllipseTestSuite) TestMinimalEllipse() {
	expected := Ellipse{
		BaseEntity: BaseEntity{
			Handle:    "ELL",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		Center:                core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		MajorAxisEnd:          core.Point{X: 2.0, Y: 5.0, Z: 7.0},
		ExtrusionDirection:    core.Point{X: 0.0, Y: 0.0, Z: 1.0},
		MinorToMajorAxisRatio: 1.0,
		StartParameter:        0.0,
		EndParameter:          360.0,
	}

	next := core.Tagger(strings.NewReader(testMinimalEllipse))
	ellipse, err := NewEllipse(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(ellipse))

	suite.False(ellipse.IsSeqEnd())
	suite.False(ellipse.HasNestedEntities())
}

func (suite *EllipseTestSuite) TestEllipseAllAttribs() {
	expected := Ellipse{
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
		Center:                core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		MajorAxisEnd:          core.Point{X: 2.0, Y: 5.0, Z: 7.0},
		ExtrusionDirection:    core.Point{X: 32.1, Y: 12.6, Z: 95.1},
		MinorToMajorAxisRatio: 1.5,
		StartParameter:        15.0,
		EndParameter:          45.0,
	}

	next := core.Tagger(strings.NewReader(testEllipseAllAttribs))
	ellipse, err := NewEllipse(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(ellipse))
}

func (suite *EllipseTestSuite) TestEllipseOff() {
	expected := Ellipse{
		BaseEntity: BaseEntity{
			On:      false,
			Color:   4,
			Visible: true,
		},
		Center:                core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		MajorAxisEnd:          core.Point{X: 2.0, Y: 5.0, Z: 7.0},
		ExtrusionDirection:    core.Point{X: 0.0, Y: 0.0, Z: 1.0},
		MinorToMajorAxisRatio: 1.0,
		StartParameter:        0.0,
		EndParameter:          360.0,
	}

	next := core.Tagger(strings.NewReader(testEllipseOff))
	ellipse, err := NewEllipse(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(ellipse))
}

func (suite *EllipseTestSuite) TestEllipseNotEqualToDifferentType() {
	suite.False(Ellipse{}.Equals(core.NewIntegerValue(0)))
}

func TestEllipseTestSuite(t *testing.T) {
	suite.Run(t, new(EllipseTestSuite))
}

const testMinimalEllipse = `  0
ELLIPSE
  5
ELL
  8
0
 10
1.1
 20
1.2
 30
1.3
 11
2.0
 21
5.0
 31
7.0
`

const testEllipseAllAttribs = `  0
ELLIPSE
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
 10
1.1
 20
1.2
 30
1.3
 11
2.0
 21
5.0
 31
7.0
210
32.1
220
12.6
230
95.1
 40
1.5
 41
15.0
 42
45.0
`

const testEllipseOff = `  0
ELLIPSE
 62
-4
 10
1.1
 20
1.2
 30
1.3
 11
2.0
 21
5.0
 31
7.0
`
