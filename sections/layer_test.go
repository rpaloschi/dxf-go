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

func TestLayer(t *testing.T) {
	next := core.Tagger(strings.NewReader(dxfLayer))
	tags := core.TagSlice(core.AllTags(next))

	layer, err := NewLayer(tags)

	assert.Nil(t, err)
	assert.Equal(t, "VIEW_PORT", layer.Name)
	assert.True(t, layer.Locked)
	assert.True(t, layer.Frozen)
	assert.True(t, layer.On)
	assert.Equal(t, 10, layer.Color)
	assert.Equal(t, "CONTINUOUS", layer.LineType)
}

func TestLayerDefaultValues(t *testing.T) {
	next := core.Tagger(strings.NewReader(""))
	tags := core.TagSlice(core.AllTags(next))

	layer, err := NewLayer(tags)

	assert.Nil(t, err)
	assert.Equal(t, "", layer.Name)
	assert.False(t, layer.Locked)
	assert.False(t, layer.Frozen)
	assert.True(t, layer.On)
	assert.Equal(t, 7, layer.Color)
	assert.Equal(t, "", layer.LineType)
}

func TestLockedLayer(t *testing.T) {
	next := core.Tagger(strings.NewReader("  70\n4"))
	tags := core.TagSlice(core.AllTags(next))

	layer, _ := NewLayer(tags)

	assert.True(t, layer.Locked)
}

func TestFrozenLayer(t *testing.T) {
	next := core.Tagger(strings.NewReader("  70\n1"))
	tags := core.TagSlice(core.AllTags(next))

	layer, _ := NewLayer(tags)

	assert.True(t, layer.Frozen)
}

func TestOffLayer(t *testing.T) {
	next := core.Tagger(strings.NewReader("  62\n-4"))
	tags := core.TagSlice(core.AllTags(next))

	layer, _ := NewLayer(tags)

	assert.False(t, layer.On)
	assert.Equal(t, 4, layer.Color)
}
