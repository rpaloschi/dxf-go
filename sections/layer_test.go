package sections

import (
	"strings"
	"testing"

	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
)

const dxfLayer = `  2
VIEW_PORT
  70
5
  62
10
  6
CONTINUOUS
`

func layerFromDxfFragment(fragment string) (*Layer, error) {
	next := core.Tagger(strings.NewReader(fragment))
	tags := core.TagSlice(core.AllTags(next))
	return NewLayer(tags)
}

func TestLayer(t *testing.T) {
	layer, err := layerFromDxfFragment(dxfLayer)

	assert.Nil(t, err)
	assert.Equal(t, "VIEW_PORT", layer.Name)
	assert.True(t, layer.Locked)
	assert.True(t, layer.Frozen)
	assert.True(t, layer.On)
	assert.Equal(t, int64(10), layer.Color)
	assert.Equal(t, "CONTINUOUS", layer.LineType)
}

func TestLayerDefaultValues(t *testing.T) {
	layer, err := layerFromDxfFragment("")

	assert.Nil(t, err)
	assert.Equal(t, "", layer.Name)
	assert.False(t, layer.Locked)
	assert.False(t, layer.Frozen)
	assert.True(t, layer.On)
	assert.Equal(t, int64(7), layer.Color)
	assert.Equal(t, "", layer.LineType)
}

func TestLockedLayer(t *testing.T) {
	layer, _ := layerFromDxfFragment("  70\n4")

	assert.True(t, layer.Locked)
}

func TestFrozenLayer(t *testing.T) {
	layer, _ := layerFromDxfFragment("  70\n1")

	assert.True(t, layer.Frozen)
}

func TestOffLayer(t *testing.T) {
	layer, _ := layerFromDxfFragment("  62\n-4")

	assert.False(t, layer.On)
	assert.Equal(t, int64(4), layer.Color)
}

const sampleLayerTable = `  0
TABLE
  2
LAYER
 70
3
  0
LAYER
  2
0
 70
0
 62
7
  6
CONTINUOUS
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
`

func TestNewLayerTable(t *testing.T) {
	expected := map[string]*Layer{
		"0": {
			Name: "0", Color: 7, LineType: "CONTINUOUS", Locked: false, Frozen: false, On: true},
		"VIEW_PORT": {
			Name: "VIEW_PORT", Color: 3, LineType: "DASHED", Locked: true, Frozen: true, On: false},
	}

	next := core.Tagger(strings.NewReader(sampleLayerTable))
	tags := core.TagSlice(core.AllTags(next))

	table, err := NewLayerTable(tags)

	assert.Nil(t, err)
	assert.Equal(t, len(expected), len(table))

	for key, expectedLayer := range expected {
		layer := table[key]

		assert.True(t, expectedLayer.Equals(layer))
	}
}

const invalidTableTags = `  0
TABLE
  2
LAYER
  0
LAYER
  20
1.1
`

func TestNewLayerTableInvalidTable(t *testing.T) {
	next := core.Tagger(strings.NewReader(invalidTableTags))
	tags := core.TagSlice(core.AllTags(next))

	_, err := NewLayerTable(tags)

	assert.Equal(t, "Invalid table. Missing TABLE AND/OR ENDTAB tags.", err.Error())
}

func TestNewLayerTableWrongTagType(t *testing.T) {
	next := core.Tagger(strings.NewReader(sampleLayerTable))
	tags := core.TagSlice(core.AllTags(next))

	tags[6].Value = core.NewStringValue("im an int ;-)")

	_, err := NewLayerTable(tags)

	assert.Equal(t,
		"Error parsing type of &core.String{value:\"im an int ;-)\"} as an Integer",
		err.Error())
}

func TestCompareLayerWrongType(t *testing.T) {
	layer, _ := layerFromDxfFragment(dxfLayer)
	assert.False(t, layer.Equals(core.NewStringValue("str")))
}
