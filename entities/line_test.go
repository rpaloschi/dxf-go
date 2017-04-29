package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type LineTestSuite struct {
	suite.Suite
}

func (suite *LineTestSuite) TestMinimalLine() {
	expected := Line{
		BaseEntity: BaseEntity{
			Handle:    "LH",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		Start:              core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		End:                core.Point{X: 2.0, Y: 5.0, Z: 7.0},
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalLine))
	line, err := NewLine(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(line))

	suite.False(line.IsSeqEnd())
	suite.False(line.HasNestedEntities())
}

func (suite *LineTestSuite) TestLineAllAttribs() {
	expected := Line{
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
		Start:              core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		End:                core.Point{X: 2.0, Y: 5.0, Z: 7.0},
		ExtrusionDirection: core.Point{X: 32.1, Y: 12.6, Z: 95.1},
	}

	next := core.Tagger(strings.NewReader(testLineAllAttribs))
	line, err := NewLine(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(line))
}

func (suite *LineTestSuite) TestLineOff() {
	expected := Line{
		BaseEntity: BaseEntity{
			On:      false,
			Color:   4,
			Visible: true,
		},
		Start:              core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		End:                core.Point{X: 2.0, Y: 5.0, Z: 7.0},
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testLineOff))
	line, err := NewLine(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(line))
}

func (suite *LineTestSuite) TestLineNotEqualToDifferentType() {
	suite.False(Line{}.Equals(core.NewIntegerValue(0)))
}

func TestLineTestSuite(t *testing.T) {
	suite.Run(t, new(LineTestSuite))
}

const testMinimalLine = `  0
LINE
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
 11
2.0
 21
5.0
 31
7.0
`

const testLineAllAttribs = `  0
LINE
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
`

const testLineOff = `  0
LINE
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
