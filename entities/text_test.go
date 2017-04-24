package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type TextTestSuite struct {
	suite.Suite
}

func (suite *TextTestSuite) TestMinimalText() {
	expected := Text{
		BaseEntity: BaseEntity{
			Handle:    "3E5",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		RelativeXScale:     1.0,
		StyleName:          "STANDARD",
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalText))
	text, err := NewText(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(text))
}

func (suite *TextTestSuite) TestTextAllAttribs() {
	expected := Text{
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
		Thickness:               55.221,
		FirstAlignmentPoint:     core.Point{X: 10.5, Y: 11.2, Z: 76.3},
		Height:                  2.5,
		Value:                   "THIS IS A TEXT",
		Rotation:                20.5,
		RelativeXScale:          2.0,
		ObliqueAngle:            10.0,
		StyleName:               "MY STYLE",
		MirroredX:               true,
		MirroredY:               true,
		HorizontalJustification: HTEXT_RIGHT,
		SecondAlignmentPoint:    core.Point{X: 11.5, Y: 12.2, Z: 77.3},
		ExtrusionDirection:      core.Point{X: 32.1, Y: 12.6, Z: 95.1},
		VerticalJustification:   VTEXT_TOP,
	}

	next := core.Tagger(strings.NewReader(testTextAllAttribs))
	text, err := NewText(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(text))
}

func (suite *TextTestSuite) TestTextNotEqualToDifferentType() {
	suite.False(Text{}.Equals(core.NewStringValue("AAA")))
}

func TestTextTestSuite(t *testing.T) {
	suite.Run(t, new(TextTestSuite))
}

const testMinimalText = `  0
TEXT
  5
3E5
  8
0
`

const testTextAllAttribs = `  0
TEXT
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
55.221
 10
10.5
 20
11.2
 30
76.3
 40
2.5
  1
THIS IS A TEXT
 50
20.5
 41
2
 51
10.0
  7
MY STYLE
 71
6
 72
2
 73
3
 11
11.5
 21
12.2
 31
77.3
210
32.1
220
12.6
230
95.1
`
