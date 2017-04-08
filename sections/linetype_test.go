package sections

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func lineTypeFromDxfFragment(fragment string) (*LineType, error) {
	next := core.Tagger(strings.NewReader(fragment))
	tags := core.TagSlice(core.AllTags(next))
	return NewLineType(tags)
}

func TestLineTypeDefaultValues(t *testing.T) {
	lineType, err := lineTypeFromDxfFragment("")

	assert.Nil(t, err)
	assert.Equal(t, "", lineType.Name)
	assert.Equal(t, "", lineType.Description)
	assert.InDelta(t, 0.0, lineType.Length, 0.001)
	assert.Equal(t, []*LineElement{}, lineType.Pattern)
}

const dxfLineType = ` 2
DASHDOT
  3
Strange dashdot
 72
    65
 73
     4
 40
1.0
 49
0.4
 74
3
 46
1.0
 50
0.5
 44
1.5
 45
2.3
  9
Sample Text
 49
0.3
 74
5
 75
1
 46
1.5
 50
1.0
 44
0.1
 45
0.2
 49
0.3
 74
4
 75
2
`

func TestParseLineType(t *testing.T) {
	lineType, err := lineTypeFromDxfFragment(dxfLineType)

	assert.Nil(t, err)
	assert.Equal(t, "DASHDOT", lineType.Name)
	assert.Equal(t, "Strange dashdot", lineType.Description)
	assert.InDelta(t, 1.0, lineType.Length, 0.001)

	expectedElements := []*LineElement{
		{0.4, true, true, false, 0,
			1.0, 0.5, 1.5, 2.3, "Sample Text"},
		{0.3, true, false, true, 1,
			1.5, 1.0, 0.1, 0.2, ""},
		{0.3, false, false, true, 2,
			1.0, 0.0, 0.0, 0.0, ""},
	}

	assert.Equal(t, expectedElements, lineType.Pattern)
}

const lineTypeTable = `  0
TABLE
  2
LTYPE
  5
21
 70
2
  0
LTYPE
  5
B
  2
CONTINUOUS
 70
0
  3
Solid line
 72
65
 73
0
 40
0.0
  0
LTYPE
  5
3E1
  2
DASHED
 70
0
  3
__ __
 72
    65
 73
     2
 40
0.75
 49
0.5
 74
4
 75
1
 49
0.25
 74
3
 46
2.2
 50
0.5
 44
1.3
 45
2.5
  9
X
  0
ENDTAB
`

func TestNewLineTypeTable(t *testing.T) {
	expected := map[string]*LineType{
		"CONTINUOUS": {
			Name:        "CONTINUOUS",
			Description: "Solid line",
			Length:      0.0,
			Pattern:     []*LineElement{}},
		"DASHED": {
			Name:        "DASHED",
			Description: "__ __",
			Length:      0.75,
			Pattern: []*LineElement{
				{
					Length:           0.5,
					AbsoluteRotation: false,
					IsTextString:     false,
					IsShape:          true,
					ShapeNumber:      1,
					Scale:            1.0,
					RotationAngle:    0.0,
					XOffset:          0.0,
					YOffset:          0.0,
					Text:             "",
				},
				{
					Length:           0.25,
					AbsoluteRotation: true,
					IsTextString:     true,
					IsShape:          false,
					ShapeNumber:      0,
					Scale:            2.2,
					RotationAngle:    0.5,
					XOffset:          1.3,
					YOffset:          2.5,
					Text:             "X",
				},
			}},
	}

	next := core.Tagger(strings.NewReader(lineTypeTable))
	tags := core.TagSlice(core.AllTags(next))

	table, err := NewLineTypeTable(tags)

	assert.Nil(t, err)
	assert.Equal(t, len(expected), len(table))

	for key, expectedLineType := range expected {
		ltype := table[key]

		assert.True(t, expectedLineType.Equals(*ltype),
			"Expected %+v and %+v to be equal",
			spew.Sdump(expectedLineType), spew.Sdump(ltype))
	}
}

const invalidLTypeTableTags = `  0
TABLE
  2
LAYER
  0
LAYER
  20
1.1
`

func TestNewLineTypeTableInvalidTable(t *testing.T) {
	next := core.Tagger(strings.NewReader(invalidLTypeTableTags))
	tags := core.TagSlice(core.AllTags(next))

	_, err := NewLineTypeTable(tags)

	assert.Equal(t, "Invalid table. Missing TABLE AND/OR ENDTAB tags.", err.Error())
}

func TestNewLineTypeTableWrongTagType(t *testing.T) {
	next := core.Tagger(strings.NewReader(lineTypeTable))
	tags := core.TagSlice(core.AllTags(next))

	tags[11].Value = core.NewStringValue("im a fake float")

	_, err := NewLineTypeTable(tags)

	assert.Equal(t,
		"Error parsing type of &core.String{value:\"im a fake float\"} as a Float",
		err.Error())
}

func TestLineTypeEquals(t *testing.T) {
	tests := []struct {
		lhs    *LineType
		rhs    *LineType
		equals bool
	}{
		{
			&LineType{Name: "A", Description: "Desc A"},
			&LineType{Name: "A", Description: "Desc A"},
			true,
		},
		{
			&LineType{Name: "A", Description: "Desc A"},
			&LineType{Name: "B", Description: "Desc A"},
			false,
		},
		{
			&LineType{Name: "A", Description: "Desc A", Pattern: []*LineElement{{Length: 0.0}}},
			&LineType{Name: "A", Description: "Desc A", Pattern: []*LineElement{{Length: 1.0}}},
			false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.lhs.Equals(*test.rhs), test.equals)
	}
}

const dxfLineTypeShapeNrNoShapeOrText = `  0
LTYPE
  5
3E1
  2
DASHED
 70
0
  3
__ __
 72
    65
 73
     1
 40
0.75
 49
0.75
 74
0
 75
1
`

func TestParseLineTypeShapeNr(t *testing.T) {
	lineType, err := lineTypeFromDxfFragment(dxfLineTypeShapeNrNoShapeOrText)

	assert.Nil(t, err)
	assert.Equal(t, 0, lineType.Pattern[0].ShapeNumber)
}

const dxfLineTypeShapeNrNot0AndText = `  0
LTYPE
  5
3E1
  2
DASHED
 70
0
  3
__ __
 72
    65
 73
     1
 40
0.75
 49
0.75
 74
2
 75
1
`

func TestParseLineTypeShapeNrWhenIsText(t *testing.T) {
	lineType, err := lineTypeFromDxfFragment(dxfLineTypeShapeNrNot0AndText)

	assert.Nil(t, err)
	assert.Equal(t, 0, lineType.Pattern[0].ShapeNumber)
}
