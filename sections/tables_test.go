package sections

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStringMappedTablesAreEquals(t *testing.T) {
	tests := []struct {
		t1     StringMappedTable
		t2     StringMappedTable
		equals bool
	}{
		{
			t1:     LayerTable{"VIEW_PORT": {Name: "VIEW_PORT"}},
			t2:     LayerTable{"VIEW_PORT": {Name: "VIEW_PORT"}},
			equals: true,
		},
		{
			t1:     StyleTable{"Style": {Name: "Style"}},
			t2:     StyleTable{"Style": {Name: "Style"}},
			equals: true,
		},
		{
			t1:     LineTypeTable{"LineTypeTable": {Name: "LineTypeTable"}},
			t2:     LineTypeTable{"LineTypeTable": {Name: "LineTypeTable"}},
			equals: true,
		},
		{
			t1: LineTypeTable{"LineTypeTable": {Name: "LineTypeTable"}},
			t2: LineTypeTable{
				"LineTypeTable":  {Name: "LineTypeTable"},
				"LineTypeTable2": {Name: "LineTypeTable2"}},
			equals: false,
		},
		{
			t1:     LayerTable{"VIEW_PORT": {Name: "VIEW_PORT"}},
			t2:     LayerTable{"OTHER": {Name: "OTHER"}},
			equals: false,
		},
		{
			t1:     LayerTable{"VIEW_PORT": {Name: "VIEW_PORT"}},
			t2:     LayerTable{"VIEW_PORT": {Name: "OTHER"}},
			equals: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, StringMappedTablesAreEquals(test.t1, test.t2), test.equals)
	}
}

func TestNewTablesSection(t *testing.T) {
	expected := TablesSection{
		Layers: LayerTable{
			"VIEW_PORT": {
				Name:     "VIEW_PORT",
				Color:    3,
				LineType: "DASHED",
				Locked:   true,
				Frozen:   true,
				On:       false,
			},
		},
		LineTypes: LineTypeTable{
			"CONTINUOUS": {
				Name:        "CONTINUOUS",
				Description: "Solid line",
				Length:      1.0,
				Pattern:     []*LineElement{},
			},
		},
		Styles: StyleTable{
			"H_TEXT": {
				Name:           "H_TEXT",
				Height:         1.0,
				Width:          1.0,
				Oblique:        0.0,
				IsBackwards:    false,
				IsUpsideDown:   false,
				IsShape:        false,
				IsVerticalText: false,
				Font:           "txt",
				BigFont:        "stxt",
			},
		},
	}

	next := core.Tagger(strings.NewReader(dxfTablesSection))
	tags := core.TagSlice(core.AllTags(next))

	tablesSection, err := NewTablesSection(tags)

	assert.Equal(t, nil, err)
	assert.True(t, expected.Equals(tablesSection),
		"Expected %+v and %+v to be equals",
		spew.Sdump(expected), spew.Sdump(tablesSection))
}

const dxfTablesSection = `  0
SECTION
  2
TABLES
  0
TABLE
  2
LAYER
 70
1
  0
LAYER
  2
VIEW_PORT
 70
5
 62
-3
  6
DASHED
  0
ENDTAB
  0
TABLE
  2
LTYPE
  5
21
 70
1
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
 40
1.0
  0
ENDTAB
  0
TABLE
  2
STYLE
 70
1
  0
STYLE
  2
H_TEXT
  3
txt
  4
stxt
  0
ENDTAB
  0
ENDSEC
`

func TestNewTablesSectionInvalidTableEntryTags(t *testing.T) {

	next := core.Tagger(strings.NewReader(dxfTablesSectionInvalidTableEntryTags))
	tags := core.TagSlice(core.AllTags(next))

	_, err := NewTablesSection(tags)

	assert.Equal(t, "Invalid table. Missing TABLE AND/OR ENDTAB tags.", err.Error())
}

const dxfTablesSectionInvalidTableEntryTags = `  0
SECTION
  2
TABLES
  0
TABLE
  2
LAYER
 70
1
  0
LAYER
  2
VIEW_PORT
  0
ENDSEC
`

func TestNewTablesSectionInvalidTag(t *testing.T) {
	next := core.Tagger(strings.NewReader(dxfTablesSectionSimple))
	tags := core.TagSlice(core.AllTags(next))

	tags[7].Value = core.NewStringValue("im an int ;-)")

	_, err := NewTablesSection(tags)

	assert.NotNil(t, err)
}

const dxfTablesSectionSimple = `  0
SECTION
  2
TABLES
  0
TABLE
  2
LAYER
 70
1
  0
LAYER
  2
VIEW_PORT
 62
1
  0
ENDTAB
  0
ENDSEC
`

func TestNewTablesSectionUnknownTableTypeDoesNotFail(t *testing.T) {
	next := core.Tagger(strings.NewReader(dxfTablesUnknownType))
	tags := core.TagSlice(core.AllTags(next))

	_, err := NewTablesSection(tags)

	assert.Nil(t, err)
}

const dxfTablesUnknownType = `  0
SECTION
  2
TABLES
  0
TABLE
  2
LAYER
 70
1
  0
LAYER
  2
VIEW_PORT
 70
5
 62
-3
  6
DASHED
  0
ENDTAB
  0
TABLE
  2
INVALID_TYPE
  5
21
 70
1
  0
INVALID_TYPE
  5
B
  2
CONTINUOUS
 70
0
  3
Solid line
 40
1.0
  0
ENDTAB
  0
TABLE
  2
STYLE
 70
1
  0
STYLE
  2
H_TEXT
  3
txt
  4
stxt
  0
ENDTAB
  0
ENDSEC
`

func TestDifferentTablesSection(t *testing.T) {
	section := TablesSection{Layers: LayerTable{"VIEW_PORT": {Name: "VIEW_PORT"}}}

	assert.Equal(t, section.Equals(core.NewIntegerValue(1)), false)
}
