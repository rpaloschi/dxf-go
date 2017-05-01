package document

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rpaloschi/dxf-go/core"
	"github.com/rpaloschi/dxf-go/entities"
	"github.com/rpaloschi/dxf-go/sections"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var expectedDocument DxfDocument

func TestDxfDocumentFromStream(t *testing.T) {
	doc, err := DxfDocumentFromStream(strings.NewReader(testSimpleDxf))

	assert.Nil(t, err)
	assert.True(t, expectedDocument.Equals(doc),
		"Expected %+v and %+v to be equals",
		spew.Sdump(expectedDocument), spew.Sdump(doc))
}

func TestDxfDocumentErrorParsingSection(t *testing.T) {
	doc, err := DxfDocumentFromStream(strings.NewReader(testWrongSection))

	assert.Nil(t, doc)
	assert.NotNil(t, err)
}

func TestDxfDocumentFromStreamWithInvalidSection(t *testing.T) {
	doc, err := DxfDocumentFromStream(
		strings.NewReader(testSimpleDxfWithInvalidSection))

	assert.Nil(t, err)
	assert.True(t, expectedDocument.Equals(doc),
		"Expected %+v and %+v to be equals",
		spew.Sdump(expectedDocument), spew.Sdump(doc))
}

const testSimpleDxf = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
AC1021
  9
$INSBASE
 10
0.1
 20
22.0
 30
53.5
  0
ENDSEC
  0
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
  0
SECTION
  2
ENTITIES
  0
POINT
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
  0
ENDSEC
  0
SECTION
  2
BLOCKS
  0
BLOCK
  8
1
  2
2
 70
64
 10
1.0
 20
3.0
  0
ENDBLK
  0
ENDSEC
`

const testWrongSection = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
AC1021
  0
ENDSEC
  0
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
ENDSEC
`

const testSimpleDxfWithInvalidSection = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
AC1021
  9
$INSBASE
 10
0.1
 20
22.0
 30
53.5
  0
ENDSEC
  0
SECTION
  2
IDONTEXIST
  0
ENDSEC
  0
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
  0
SECTION
  2
ENTITIES
  0
POINT
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
  0
ENDSEC
  0
SECTION
  2
BLOCKS
  0
BLOCK
  8
1
  2
2
 70
64
 10
1.0
 20
3.0
  0
ENDBLK
  0
ENDSEC
`

func init() {
	expectedDocument = DxfDocument{
		Header: &sections.HeaderSection{
			Values: map[string]core.TagSlice{
				"$ACADVER": {
					core.NewTag(1, core.NewStringValue("AC1021")),
				},
				"$INSBASE": {
					core.NewTag(10, core.NewFloatValue(0.1)),
					core.NewTag(20, core.NewFloatValue(22.0)),
					core.NewTag(30, core.NewFloatValue(53.5)),
				},
				"$DWGCODEPAGE": {
					core.NewTag(3, core.NewStringValue("ANSI_1252")),
				},
			},
		},
		Tables: &sections.TablesSection{
			Layers: sections.Table{
				"VIEW_PORT": &sections.Layer{
					Name:  "VIEW_PORT",
					Color: 1,
					On:    true,
				},
			},
		},
		Entities: &sections.EntitiesSection{
			Entities: entities.EntitySlice{
				&entities.Point{
					BaseEntity: entities.BaseEntity{
						Handle:    "3E5",
						On:        true,
						Visible:   true,
						LayerName: "0",
					},
					Location:           core.Point{X: 1.1, Y: 1.2, Z: 1.3},
					ExtrusionDirection: core.Point{X: 0, Y: 0, Z: 1},
				},
			},
		},
		Blocks: sections.BlocksSection{
			"2": &sections.Block{
				Name:      "2",
				LayerName: "1",
				BasePoint: core.Point{X: 1, Y: 3},
			},
		},
	}
}
