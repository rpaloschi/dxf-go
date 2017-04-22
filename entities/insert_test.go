package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type InsertTestSuite struct {
	suite.Suite
}

func (suite *InsertTestSuite) TestMinimalInsert() {
	expected := Insert{
		BaseEntity: BaseEntity{
			Handle:    "INS",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		InsertionPoint:     core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		ScaleFactorX:       1.0,
		ScaleFactorY:       1.0,
		ScaleFactorZ:       1.0,
		RotationAngle:      0.0,
		ColumnCount:        1,
		RowCount:           1,
		ColumnSpacing:      0.0,
		RowSpacing:         0.0,
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalInsert))
	arc, err := NewInsert(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *InsertTestSuite) TestInsertAllAttribs() {
	expected := Insert{
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
		BlockName:          "B1",
		InsertionPoint:     core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		ScaleFactorX:       2.2,
		ScaleFactorY:       5.7,
		ScaleFactorZ:       9.6,
		RotationAngle:      45.0,
		ColumnCount:        2,
		RowCount:           3,
		ColumnSpacing:      6.0,
		RowSpacing:         7.0,
		ExtrusionDirection: core.Point{X: 32.1, Y: 12.6, Z: 95.1},
	}

	next := core.Tagger(strings.NewReader(testInsertAllAttribs))
	arc, err := NewInsert(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *InsertTestSuite) TestInsertOff() {
	expected := Insert{
		BaseEntity: BaseEntity{
			On:      false,
			Color:   4,
			Visible: true,
		},
		InsertionPoint:     core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		ScaleFactorX:       1.0,
		ScaleFactorY:       1.0,
		ScaleFactorZ:       1.0,
		RotationAngle:      0.0,
		ColumnCount:        1,
		RowCount:           1,
		ColumnSpacing:      0.0,
		RowSpacing:         0.0,
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testInsertOff))
	arc, err := NewInsert(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func (suite *InsertTestSuite) TestInsertNotEqualToDifferentType() {
	suite.False(Insert{}.Equals(core.NewIntegerValue(0)))
}

func TestInsertTestSuite(t *testing.T) {
	suite.Run(t, new(InsertTestSuite))
}

const testMinimalInsert = `  0
INSERT
  5
INS
  8
0
 10
1.1
 20
1.2
 30
1.3
`

const testInsertAllAttribs = `  0
INSERT
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
  2
B1
 41
2.2
 42
5.7
 43
9.6
 44
6
 45
7
 50
45
 70
2
 71
3
210
32.1
220
12.6
230
95.1
`

const testInsertOff = `  0
INSERT
 62
-4
 10
1.1
 20
1.2
 30
1.3
`
