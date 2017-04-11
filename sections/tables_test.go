package sections

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewTablesSection(t *testing.T) {
	expected := TablesSection{
		LayerTable: map[string]*Layer{
			"VIEW_PORT": {
				Name:     "VIEW_PORT",
				Color:    3,
				LineType: "DASHED",
				Locked:   true,
				Frozen:   true,
				On:       false,
			},
		},
		LineTypeTable: map[string]*LineType{
			"CONTINUOUS": {
				Name:        "CONTINUOUS",
				Description: "Solid line",
				Length:      1.0,
				Pattern:     []*LineElement{},
			},
		},
		StyleTable: map[string]*Style{
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
	assert.True(t, expected.Equals(*tablesSection),
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
