package sections

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
	layer, err := NewLayer(tags)
	return layer, err
}

func TestLayer(t *testing.T) {
	layer, err := layerFromDxfFragment(dxfLayer)

	assert.Nil(t, err)
	assert.Equal(t, "VIEW_PORT", layer.Name)
	assert.True(t, layer.Locked)
	assert.True(t, layer.Frozen)
	assert.True(t, layer.On)
	assert.Equal(t, 10, layer.Color)
	assert.Equal(t, "CONTINUOUS", layer.LineType)
}

func TestLayerDefaultValues(t *testing.T) {
	layer, err := layerFromDxfFragment("")

	assert.Nil(t, err)
	assert.Equal(t, "", layer.Name)
	assert.False(t, layer.Locked)
	assert.False(t, layer.Frozen)
	assert.True(t, layer.On)
	assert.Equal(t, 7, layer.Color)
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
	assert.Equal(t, 4, layer.Color)
}
