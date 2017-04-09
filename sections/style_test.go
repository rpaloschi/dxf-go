package sections

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func styleFromDxfFragment(fragment string) (*Style, error) {
	next := core.Tagger(strings.NewReader(fragment))
	tags := core.TagSlice(core.AllTags(next))
	return NewStyle(tags)
}

func TestStyleDefaultValues(t *testing.T) {
	style, err := styleFromDxfFragment("")

	assert.Nil(t, err)
	assert.Equal(t, "", style.Name)
	assert.InDelta(t, 1.0, style.Height, 0.001)
	assert.InDelta(t, 1.0, style.Width, 0.001)
	assert.InDelta(t, 0.0, style.Oblique, 0.001)
	assert.False(t, style.IsBackwards)
	assert.False(t, style.IsUpsideDown)
	assert.False(t, style.IsShape)
	assert.False(t, style.IsVerticalText)
	assert.Equal(t, "", style.Font)
	assert.Equal(t, "", style.BigFont)
}

const dxfStyle = `  2
STANDARD
 70
     5
 40
3.55
 41
1.1
 50
6.0
 71
     6
 42
0.2
  3
txt
  4
Arial.ttf
`

func TestDxfStyle(t *testing.T) {
	style, err := styleFromDxfFragment(dxfStyle)

	assert.Nil(t, err)
	assert.Equal(t, "STANDARD", style.Name)
	assert.InDelta(t, 3.55, style.Height, 0.001)
	assert.InDelta(t, 1.1, style.Width, 0.001)
	assert.InDelta(t, 6.0, style.Oblique, 0.001)
	assert.True(t, style.IsBackwards)
	assert.True(t, style.IsUpsideDown)
	assert.True(t, style.IsShape)
	assert.True(t, style.IsVerticalText)
	assert.Equal(t, "txt", style.Font)
	assert.Equal(t, "Arial.ttf", style.BigFont)
}

const dxfStyleTable = `  0
TABLE
  2
STYLE
 70
2
  0
STYLE
  2
H_TEXT
  3
txt
  4
stxt
  0
STYLE
  2
V_TEXT
 70
4
 40
1.5
 41
2.3
 50
0.5
 71
6
  3
1
  4
2
  0
STYLE
  2
SHAPE
 70
1
 50
0.0
 71
0
  0
ENDTAB
`

func TestNewStyleTable(t *testing.T) {
	expected := map[string]*Style{
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
		"V_TEXT": {
			Name:           "V_TEXT",
			Height:         1.5,
			Width:          2.3,
			Oblique:        0.5,
			IsBackwards:    true,
			IsUpsideDown:   true,
			IsShape:        false,
			IsVerticalText: true,
			Font:           "1",
			BigFont:        "2",
		},
		"SHAPE": {
			Name:           "SHAPE",
			Height:         1.0,
			Width:          1.0,
			Oblique:        0.0,
			IsBackwards:    false,
			IsUpsideDown:   false,
			IsShape:        true,
			IsVerticalText: false,
			Font:           "",
			BigFont:        "",
		},
	}

	next := core.Tagger(strings.NewReader(dxfStyleTable))
	tags := core.TagSlice(core.AllTags(next))

	table, err := NewStyleTable(tags)

	assert.Nil(t, err)
	assert.Equal(t, len(expected), len(table))

	for key, expectedStyle := range expected {
		style := table[key]

		assert.True(t, expectedStyle.Equals(*style))
	}
}

const invalidStyleTable = `  0
TABLE
  2
STYLE
  0
STYLE
  20
1.1
`

func TestNewStyleTableInvalidTable(t *testing.T) {
	next := core.Tagger(strings.NewReader(invalidStyleTable))
	tags := core.TagSlice(core.AllTags(next))

	_, err := NewStyleTable(tags)

	assert.Equal(t, "Invalid table. Missing TABLE AND/OR ENDTAB tags.", err.Error())
}

func TestNewStyleTableWrongTagType(t *testing.T) {
	next := core.Tagger(strings.NewReader(dxfStyleTable))
	tags := core.TagSlice(core.AllTags(next))

	tags[9].Value = core.NewStringValue("im a fake int")

	_, err := NewStyleTable(tags)

	assert.Equal(t,
		"Error parsing type of &core.String{value:\"im a fake int\"} as an Integer",
		err.Error())
}
